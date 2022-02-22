## Hace todos los puntos pedidos (40%)

#### Permite crear usuarios vía endpoint

- No 😔

#### Permite añadir favoritos vía endpoint

OK

#### Si el pokemon ya está marcado como favorito, ¿se lanza una excepción de dominio?

OK

#### Si el usuario no existe, ¿se lanza una excepción de dominio?

- No

#### Si se lanza una excepción desde el dominio, ¿se traduce en infraestructura a un código HTTP?

OK

#### Hay tests unitarios

- Si, pero estos deberían estar separados del `src`. Es decir, por un lado debería estar el código de la aplicación
  en `src` y a la altura de esta, debería haber otra carpeta `tests``

**Error principal: No cumple los requisitos pedidos (no se pueden crear usuarios vía endpoint)**
**Puntuación: 20/40**

## Se aplican conceptos explicados (50%)

#### Separación correcta de capas (application, domain, infrastructure + BC/module/layer)

- El servicio de aplicación `AddFavoritePokemon` tiene dos responsabilidades: crear un usuario y añadir un pokemon
  favorito. Estamos incumpliendo el SRP!

#### Aggregates + VOs

- Tendría más sentido que se llamara `PokemonFavourites` que `PokemonIdCollection`. Ya que el listado de entidades que
  forma parte de un usuario son sus pokemon favoritos.

#### No se trabajan con tipos primitivos en dominio

OK

#### Hay servicios de dominio

- Se llama directamente al repository.

#### Hay use cases en aplicación reutilizables

OK

#### Se aplica el patrón repositorio

- Hay dos patrones repositorio: `UserRepository` y `FavouritePokemonRepository` si el agregado es usuario, el único
  puerto de repositorio debería ser `UserRepository` (recordad: 1 módulo -> 1 repositorio). Aunque veo que el
  repositorio `FavoritePokemonRepository` no se está usando.

#### Se utilizan object mothers

No

**Error principal: Dos repositorios en el mismo módulo**
**Puntuación: 25/50**

## Facilidad setup + README (10%)

#### El README contiene al menos los apartados "cómo ejecutar la aplicación", "cómo user la aplicación"

- El swagger no tiene los nuevos endpoints, ni el README 😔.

#### Es sencillo seguir el apartado "cómo ejecutar la aplicación"

OK

**Puntuación: 7/10**

## Observaciones

- Estáis modificando pokemon favoritos de un usuario, la URL debería ser sobre usuario no sobre favoritos.

**PUNTUACIÓN FINAL: 52/100**
