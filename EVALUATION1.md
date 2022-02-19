# Evaluación reto 1

## Hace todos los puntos pedidos (40%)

#### Dado un nombre vía argumento, devolver sus tipos

OK

#### Dado un nombre vía endpoint, devolver sus tipos

OK

#### Si no existe el pokemon, ¿se lanza una excepción de dominio?

OK

#### Si la api da timeout, ¿se lanza una excepción de dominio?

OK

#### Si se lanza una excepción desde el dominio, ¿se traduce en infraestructura a un código HTTP/un error legible en consola?

- Por HTTP: Devuelve un mensaje de error pero da un 200. Debería mapearse a un código de respuesta HTTP con sentido para
  ese caso. Ejemplo: 404 not found

**Error principal: No sólo es importante el mensaje, sino también los códigos de respuesta**

**Puntuación final: 35/40**

## Se aplican conceptos explicados (40%)

#### Separación correcta de capas (application, domain, infrastructure + BC/module/layer)

OK

#### Aggregates + VOs

- Existe un VO `PokemonName` que no está en el agregado. Únicamente se está utilizando para hacer la búsqueda en el
  repositorio pero no pertenece a ningún agregado, ¿cómo es posible buscar por un VO que no pertenece a ningún agregado?

#### No se trabajan con tipos primitivos en dominio

OK

#### Hay servicios de dominio

- Se llama directamente al repository en vez de usar un searcher (servicio de dominio)

#### Hay use cases en aplicación reutilizables

OK

#### Se aplica el patrón repositorio

OK

**Error principal: Hay VOs que no pertenecen al agregado**

**Puntuación final: 30/40**

## Facilidad setup + README (20%)

#### El README contiene al menos los apartados "cómo ejecutar la aplicación", "cómo usar la aplicación"

OK

#### Es sencillo seguir el apartado "cómo ejecutar la aplicación"

OK

**Puntuación final: 20/20**

## Extra

- Docker
- Commits en "baby steps". Pequeños y legibles

**Puntuación: +8**

## Observaciones

- El módulo dentro del BC `Pokemon` se sobreentiende que son tipos de pokemons, no es necesario prefijarlo 😉
- No commiteeis/entreguéis código que no se usa (`NOT_FOUND`, `BAD_REQUEST` en `Get.go`)
- Usad
  el [estándar](https://www.theserverside.com/blog/Coffee-Talk-Java-News-Stories-and-Opinions/Why-you-should-make-kebab-case-a-URL-naming-convention-best-practice)
  de escritura de URLs (kebab case). Ej: `/pokemon-type` o `/type` pasándole el `name` como un query param
- Aunque el lenguaje os permita hacer un return de múltiples objetos, intentad encapsular la información a retornar en
  un DTO o algo similar para hacer más programación orientada a objetos :)

**PUNTUACIÓN FINAL: 93/100**
