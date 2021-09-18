package wikidataReader

func ExtractClaimsFromEntity(e *Entity, propertyName string) []*EntityClaim {
	props := make([]*EntityClaim, 0)

	for _, claimSlice := range e.Claims {
		for i := range claimSlice {
			claim := claimSlice[i]
			if claim.Mainsnak.Property == propertyName {
				props = append(props, &claim)
			}
		}
	}

	return props
}
