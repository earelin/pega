/*
 * This program is free software: you can redistribute it and/or modify it under
 * the terms of the GNU General Public License as published by the Free Software
 * Foundation, either version 3 of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY
 * WARRANTY; without even the implied warranty of MERCHANTABILITY or
 * FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License
 * for more details.
 *
 * You should have received a copy of the GNU General Public License along with
 * this program. If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/earelin/pega/tools/inebase/pkg/config"
	"github.com/earelin/pega/tools/inebase/pkg/importer"
	"github.com/earelin/pega/tools/inebase/pkg/repository"
	"io"
	"os"
	"strings"
)

func main() {
	start(os.Stdout, os.Args)
}

func start(w io.Writer, args []string) {
	var conf config.Config
	var err error

	conf, err = parseArgs(w, args)
	if errors.Is(err, flag.ErrHelp) {
		os.Exit(0)
	}
	if err != nil {
		fmt.Fprintln(w, "Erro arrancando o programa: ", err)
		showUsage()
		os.Exit(1)
	}

	err = validateConfiguration(conf)
	if err != nil {
		fmt.Fprintln(w, "Erro:", err)
	}

	var r *repository.Repository
	r, err = repository.NewRepository(conf.RepositoryConfig, context.Background())
	if err != nil {
		fmt.Fprintf(w, "Error configurado a conexión á base de datos: %s", err)
		os.Exit(1)
	}
	defer func(r *repository.Repository) {
		err := r.CloseConnection()
		if err != nil {
			fmt.Fprintf(w, "Erro pechando a conexión á base de datos: %s", err)
			os.Exit(1)
		}
	}(r)

	err = importer.ImportarConcellos(conf, r)
	if err != nil {
		fmt.Fprintf(w, "Erro importando %s: %s", conf.DataSet, err)
		os.Exit(1)
	}
}

func parseArgs(w io.Writer, args []string) (config.Config, error) {
	var c config.Config

	fs := flag.NewFlagSet("inebase", flag.ContinueOnError)
	fs.SetOutput(w)

	fs.StringVar(&c.RepositoryConfig.Filename, "file",
		"database.sqlite", "Ficheiro da base de datos")

	err := fs.Parse(args[1:])
	if err != nil {
		return c, err
	}

	if fs.NArg() != 2 {
		return c, errors.New("nome de ficheiro e dataset non expecificado")
	}

	var dataSet = strings.TrimSpace(fs.Arg(0))
	if dataSet == "" {
		return c, errors.New("nome de dataset non expecificado")
	}
	c.DataSet = dataSet

	var filePath = strings.TrimSpace(fs.Arg(1))
	if filePath == "" {
		return c, errors.New("nome de ficherio non expecificado")
	}
	c.FilePath = filePath

	return c, nil
}

func showUsage() {
	fmt.Println(`	Uso:
		inebase [opcions] dataset ficheiro

		dataset: Tipo de conxunto de datos (concellos)
		ficheiro: Ruta ao ficheiro

	Opcións:
		-h/--help Show this help info`)
}

func validateConfiguration(conf config.Config) error {
	var err error

	_, err = os.Stat(conf.FilePath)
	if err != nil {
		return fmt.Errorf("non se pode abrir o ficheiro: %w", err)
	}

	if conf.DataSet != "concellos" {
		return fmt.Errorf("dataset descoñecido: %s", conf.DataSet)
	}

	return err
}
