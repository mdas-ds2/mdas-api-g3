## Hace todos los puntos pedidos (40%)

#### Permite crear usuarios v铆a endpoint

- No 馃様

#### Permite a帽adir favoritos v铆a endpoint

OK

#### Si el pokemon ya est谩 marcado como favorito, 驴se lanza una excepci贸n de dominio?

OK

#### Si el usuario no existe, 驴se lanza una excepci贸n de dominio?

- No

#### Si se lanza una excepci贸n desde el dominio, 驴se traduce en infraestructura a un c贸digo HTTP?

OK

#### Hay tests unitarios

- Si, pero estos deber铆an estar separados del `src`. Es decir, por un lado deber铆a estar el c贸digo de la aplicaci贸n
  en `src` y a la altura de esta, deber铆a haber otra carpeta `tests``

**Error principal: No cumple los requisitos pedidos (no se pueden crear usuarios v铆a endpoint)**
**Puntuaci贸n: 20/40**

## Se aplican conceptos explicados (50%)

#### Separaci贸n correcta de capas (application, domain, infrastructure + BC/module/layer)

OK

#### Aggregates + VOs

- Tendr铆a m谩s sentido que se llamara `PokemonFavourites` que `PokemonIdCollection`. Ya que el listado de entidades que
  forma parte de un usuario son sus pokemon favoritos.

#### No se trabajan con tipos primitivos en dominio

OK

#### Hay servicios de dominio

- Se llama directamente al repository.

#### Hay use cases en aplicaci贸n reutilizables

OK

#### Se aplica el patr贸n repositorio

- Hay dos patrones repositorio: `UserRepository` y `FavouritePokemonRepository` si el agregado es usuario, el 煤nico
  puerto de repositorio deber铆a ser `UserRepository` (recordad: 1 m贸dulo -> 1 repositorio). Aunque veo que el
  repositorio `FavoritePokemonRepository` no se est谩 usando.

#### Se utilizan object mothers

No

**Error principal: No hay servicio de dominio**
**Puntuaci贸n: 40/50**

## Facilidad setup + README (10%)

#### El README contiene al menos los apartados "c贸mo ejecutar la aplicaci贸n", "c贸mo user la aplicaci贸n"

- El swagger no tiene los nuevos endpoints, ni el README 馃様.

#### Es sencillo seguir el apartado "c贸mo ejecutar la aplicaci贸n"

OK

**Puntuaci贸n: 7/10**

## Observaciones

- Est谩is modificando pokemon favoritos de un usuario, la URL deber铆a ser sobre usuario no sobre favoritos.

**PUNTUACI脫N FINAL: 67/100**
