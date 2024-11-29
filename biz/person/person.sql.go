package person

var CreatePersonQuery = `
	CREATE (p:Person {name: $name, age: $age, languages: $languages})-[l:LIKES]->(t:Technology {name: $tech})
`

var GetPersonQuery = `
	MATCH (p:Person)
	WHERE p.name = $name
	RETURN p
`

var GetPersonWithTechQuery = `
	MATCH(p:Person)-[l:LIKES]->(t:Technology)
	WHERE p.name = $name
	RETURN p, COLLECT(t) as techs
`
