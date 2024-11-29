package person

var CreatePersonQuery = `
	MERGE (p:Person {name: $name, age: $age, languages: $languages})-[l:LIKES]->(t:Technology {name: $tech})
`

var GetPersonQueryByName = `
	MATCH (p:Person)
	WHERE p.name = $name
	RETURN p
`

var GetPersonWithTechQueryByName = `
	MATCH(p:Person)-[l:LIKES]->(t:Technology)
	WHERE p.name = $name
	RETURN p, COLLECT(t) as techs
`
