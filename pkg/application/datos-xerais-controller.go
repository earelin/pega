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

package application

import (
	"github.com/earelin/pega/pkg/domain"
	"github.com/gin-gonic/gin"
)

func NewDatosXeraisController(e *gin.Engine, datosXeraisRepository domain.DatosXeraisRepository) {
	c := &DatosXeraisController{repository: datosXeraisRepository}
	e.GET("/proceso-electoral/:id/datos-xerais", c.GetDatosXerais)
	e.GET(
		"/proceso-electoral/:id/datos-xerais/comunidade-autonoma/:comunidadeAutonomaId",
		c.GetDatosXeraisComunidadeAutonoma)
	e.GET("/proceso-electoral/:id/datos-xerais/provincia/:provinciaId",
		c.GetDatosXeraisProvincia)
	e.GET("/proceso-electoral/:id/datos-xerais/concello/:concelloId",
		c.GetDatosXeraisConcello)
	e.GET("/proceso-electoral/:id/datos-xerais/concello/:concelloId/:distritoId",
		c.GetDatosXeraisDistrito)
	e.GET(
		"/proceso-electoral/:id/datos-xerais/concello/:concelloId/:distritoId/:seccionId",
		c.GetDatosXeraisSeccion)
	e.GET(
		"/proceso-electoral/:id/datos-xerais/concello/:concelloId/:distritoId/:seccionId/:codigoMesa",
		c.GetDatosXeraisMesa)
}

type DatosXeraisController struct {
	repository domain.DatosXeraisRepository
}

func (c DatosXeraisController) GetDatosXerais(gc *gin.Context) {
	var uriParams struct {
		Id int `uri:"id"`
	}
	if err := gc.ShouldBindUri(&uriParams); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	ps, ok := c.repository.FindByProceso(uriParams.Id)

	if ok {
		gc.JSON(200, ps)
	} else {
		gc.Status(404)
	}
}

func (c DatosXeraisController) GetDatosXeraisComunidadeAutonoma(gc *gin.Context) {
	var uriParams struct {
		Id                   int `uri:"id"`
		ComunidadeAutonomaId int `uri:"comunidadeAutonomaId"`
	}
	if err := gc.ShouldBindUri(&uriParams); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	ps, ok := c.repository.FindByComunidadeAutonoma(uriParams.Id, uriParams.ComunidadeAutonomaId)

	if ok {
		gc.JSON(200, ps)
	} else {
		gc.Status(404)
	}
}

func (c DatosXeraisController) GetDatosXeraisProvincia(gc *gin.Context) {
	var uriParams struct {
		Id          int `uri:"id"`
		ProvinciaId int `uri:"provinciaId"`
	}
	if err := gc.ShouldBindUri(&uriParams); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	ps, ok := c.repository.FindByProvincia(uriParams.Id, uriParams.ProvinciaId)

	if ok {
		gc.JSON(200, ps)
	} else {
		gc.Status(404)
	}
}

func (c DatosXeraisController) GetDatosXeraisConcello(gc *gin.Context) {
	var uriParams struct {
		Id         int `uri:"id"`
		ConcelloId int `uri:"concelloId"`
	}
	if err := gc.ShouldBindUri(&uriParams); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	ps, ok := c.repository.FindByConcello(uriParams.Id, uriParams.ConcelloId)

	if ok {
		gc.JSON(200, ps)
	} else {
		gc.Status(404)
	}
}

func (c DatosXeraisController) GetDatosXeraisDistrito(gc *gin.Context) {
	var uriParams struct {
		Id         int `uri:"id"`
		ConcelloId int `uri:"concelloId"`
		DistritoId int `uri:"distritoId"`
	}
	if err := gc.ShouldBindUri(&uriParams); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	ps, ok := c.repository.FindByDistrito(uriParams.Id, uriParams.ConcelloId, uriParams.DistritoId)

	if ok {
		gc.JSON(200, ps)
	} else {
		gc.Status(404)
	}
}

func (c DatosXeraisController) GetDatosXeraisSeccion(gc *gin.Context) {
	var uriParams struct {
		Id         int `uri:"id"`
		ConcelloId int `uri:"concelloId"`
		DistritoId int `uri:"distritoId"`
		SeccionId  int `uri:"seccionId"`
	}
	if err := gc.ShouldBindUri(&uriParams); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	ps, ok := c.repository.FindBySeccion(uriParams.Id, uriParams.ConcelloId, uriParams.DistritoId,
		uriParams.SeccionId)

	if ok {
		gc.JSON(200, ps)
	} else {
		gc.Status(404)
	}
}

func (c DatosXeraisController) GetDatosXeraisMesa(gc *gin.Context) {
	var uriParams struct {
		Id         int    `uri:"id"`
		ConcelloId int    `uri:"concelloId"`
		DistritoId int    `uri:"distritoId"`
		SeccionId  int    `uri:"seccionId"`
		CodigoMesa string `uri:"codigoMesa"`
	}
	if err := gc.ShouldBindUri(&uriParams); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	ps, ok := c.repository.FindByMesa(uriParams.Id, uriParams.ConcelloId, uriParams.DistritoId,
		uriParams.SeccionId, uriParams.CodigoMesa)

	if ok {
		gc.JSON(200, ps)
	} else {
		gc.Status(404)
	}
}
