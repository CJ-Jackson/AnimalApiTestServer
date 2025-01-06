# API Doc
## animals

### DELETE: `/animal/delete` Delete Animal

Delete an animal

#### payload `application/json; charset=utf-8`

```json
{
  "Id": 8
}
```

#### payload `application/xml; charset=utf-8`

```xml
<?xml version="1.0" encoding="utf-8"?>
<DeleteAnimalData>
    <Id>8</Id>
</DeleteAnimalData>
```

#### response `application/json; charset=utf-8`

```json
{
  "Message": "'8' has been deleted."
}
```

#### response `application/xml; charset=utf-8`

```xml
<?xml version="1.0" encoding="utf-8"?>
<Message>
    <Message>'8' has been deleted.</Message>
</Message>
```

### GET: `/animal/get/{id}` Get animal by id

Get an animal by id number

#### arg type

```
id: int
```

#### response `application/json; charset=utf-8`

```json
{
  "Animal": {
    "Id": 8,
    "Name": "Pikachu",
    "Description": "Ash's companion"
  }
}
```

#### response `application/xml; charset=utf-8`

```xml
<?xml version="1.0" encoding="utf-8"?>
<SingleAnimal>
    <Animal Id="8">
        <Name>Pikachu</Name>
        <Description>Ash's companion</Description>
    </Animal>
</SingleAnimal>
```

### GET: `/animal/list` List animals

Get a list of all animals

#### response `application/json; charset=utf-8`

```json
{
  "Animals": [
    {
      "Id": 1,
      "Name": "Cat",
      "Description": "House Cat"
    },
    ...
  ]
}
```

#### response `application/xml; charset=utf-8`

```xml
<?xml version="1.0" encoding="utf-8"?>
<AnimalCollection>
    <Animals Id="1">
        <Name>Cat</Name>
        <Description>House Cat</Description>
    </Animals>
    ...
</AnimalCollection>
```

### POST: `/animal/post` Post animals

Add a new animal

#### payload `application/json; charset=utf-8`

```json
{
  "Animal": {
    "Id": 20,
    "Name": "Charizard",
    "Description": "Ash's companion"
  }
}
```

#### payload `application/xml; charset=utf-8`

```xml
<?xml version="1.0" encoding="utf-8"?>
<SingleAnimal>
    <Animal Id="20">
        <Name>Charizard</Name>
        <Description>Ash's companion</Description>
    </Animal>
</SingleAnimal>
```

#### response `application/json; charset=utf-8`

```json
{
  "Animal": {
    "Id": 20,
    "Name": "Charizard",
    "Description": "Ash's companion"
  }
}
```

#### response `application/xml; charset=utf-8`

```xml
<?xml version="1.0" encoding="utf-8"?>
<SingleAnimal>
    <Animal Id="20">
        <Name>Charizard</Name>
        <Description>Ash's companion</Description>
    </Animal>
</SingleAnimal>
```

### PATCH: `/animal/update/{id}` Update Animal

Updated an existing animal

#### arg type

```
id: int
```

#### payload `application/json; charset=utf-8`

```json
{
  "Animal": {
    "Id": 8,
    "Name": "Charizard",
    "Description": "Ash's companion"
  }
}
```

#### payload `application/xml; charset=utf-8`

```xml
<?xml version="1.0" encoding="utf-8"?>
<SingleAnimal>
    <Animal Id="8">
        <Name>Charizard</Name>
        <Description>Ash's companion</Description>
    </Animal>
</SingleAnimal>
```

#### response `application/json; charset=utf-8`

```json
{
  "Message": "'8' has been updated."
}
```

#### response `application/xml; charset=utf-8`

```xml
<?xml version="1.0" encoding="utf-8"?>
<Message>
    <Message>'8' has been updated.</Message>
</Message>
```
