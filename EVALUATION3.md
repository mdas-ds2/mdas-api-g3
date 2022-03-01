## Hace todos los puntos pedidos (40%)

#### Permite obtener los detalles de un pokemon vía endpoint

OK

#### Si no existe el pokemon, ¿se lanza una excepción de dominio?

OK

#### Si la api da timeout, ¿se lanza una excepción de dominio?

OK

#### Si se lanza una excepción desde el dominio, ¿se traduce en infraestructura a un código HTTP?

OK

#### Tests de aceptación

- OK, aunque llamar a los tests de aceptación XXXControllerTest no es el nombre más correcto, ya que así parecen tests
  de integración (no estamos testeando únicamente el controller).

#### Tests de integración

OK

**Puntuación: 38/40**

## Se aplican conceptos explicados (50%)

#### Separación correcta de capas (application, domain, infrastructure + BC/module/layer)

OK

#### Aggregates + VOs

OK

#### No se trabajan con tipos primitivos en dominio

OK

#### Hay servicios de dominio

- Se llama directamente al repository.

#### Hay use cases en aplicación reutilizables

OK

#### Se aplica el patrón repositorio

OK

#### Se utilizan object mothers

No

**Puntuación: 37/50**

## Facilidad setup + README (10%)

#### El README contiene al menos los apartados ""cómo ejecutar la aplicación"", ""cómo usar la aplicación""

OK

#### Es sencillo seguir el apartado ""cómo ejecutar la aplicación""

OK

**Puntuación: 10/10**

## Observaciones

- La URL debería ser `http://localhost:5001/pokemon/25`, no `http://localhost:5001/pokemon/?id=25` ya que es un acceso
  por ID, no una búsqueda por parámetro.

**PUNTUACIÓN FINAL: 85/100**
