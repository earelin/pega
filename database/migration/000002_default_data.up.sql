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

INSERT INTO comunidade_autonoma(id, nome)
VALUES (1, 'Andalucía'),
       (2, 'Aragón'),
       (3, 'Principado de Asturias'),
       (4, 'Illes Balears'),
       (5, 'Canarias'),
       (6, 'Cantabria'),
       (7, 'Castilla y León'),
       (8, 'Castilla - La Mancha'),
       (9, 'Catalunya'),
       (10, 'Comunitat Valenciana'),
       (11, 'Extremadura'),
       (12, 'Galiza'),
       (13, 'Comunidad de Madrid'),
       (14, 'Región de Murcia'),
       (15, 'Comunidad Foral de Navarra'),
       (16, 'Euskadi'),
       (17, 'La Rioja'),
       (18, 'Ceuta'),
       (19, 'Melilla');

INSERT INTO provincia(id, nome, comunidade_autonoma_id)
VALUES (4, 'Almería', 1),
       (11, 'Cádiz', 1),
       (14, 'Córdoba', 1),
       (18, 'Granada', 1),
       (21, 'Huelva', 1),
       (23, 'Jaén', 1),
       (29, 'Málaga', 1),
       (41, 'Sevilla', 1),
       (22, 'Huesca', 2),
       (44, 'Teruel', 2),
       (50, 'Zaragoza', 2),
       (33, 'Asturias', 3),
       (7, 'Illes Balears', 4),
       (35, 'Las Palmas', 5),
       (38, 'Santa Cruz de Tenerife', 5),
       (39, 'Cantabria', 6),
       (5, 'Ávila', 7),
       (9, 'Burgos', 7),
       (24, 'León', 7),
       (34, 'Palencia', 7),
       (37, 'Salamanca', 7),
       (40, 'Segovia', 7),
       (42, 'Soria', 7),
       (47, 'Valladolid', 7),
       (49, 'Zamora', 7),
       (2, 'Albacete', 8),
       (13, 'Ciudad Real', 8),
       (16, 'Cuenca', 8),
       (19, 'Guadalajara', 8),
       (45, 'Toledo', 8),
       (8, 'Barcelona', 9),
       (17, 'Girona', 9),
       (25, 'Lleida', 9),
       (43, 'Tarragona', 9),
       (3, 'Alacant', 10),
       (12, 'Castelló', 10),
       (46, 'València', 10),
       (6, 'Badajoz', 11),
       (10, 'Cáceres', 11),
       (15, 'A Coruña', 12),
       (27, 'Lugo', 12),
       (32, 'Ourense', 12),
       (36, 'Pontevedra', 12),
       (28, 'Madrid', 13),
       (30, 'Murcia', 14),
       (31, 'Navarra', 15),
       (1, 'Araba', 16),
       (48, 'Bizkaia', 16),
       (20, 'Gipuzkoa', 16),
       (26, 'La Rioja', 17),
       (51, 'Ceuta', 18),
       (52, 'Melilla', 19);
