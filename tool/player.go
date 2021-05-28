package tool

type Application struct {
	ApplicationID string `json:"applicationId"`
	Version       string `json:"version"`
}
type Profile struct {
	Applications []Application `json:"applications"`
}

type Player struct {
	Profile Profile `json:"profile"`
}

type Players []Player

//it runs somehow upon decoding. Remove reflections for now.
/* func (p *Player) UnmarshalJSON(data []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	v := reflect.ValueOf(p).Elem()

	t := v.Type()
	//fmt.Println(t.Field(0))
	fmt.Println(m)
	fmt.Println(t.Field(0))
	//fmt.Println(t.Field(2))
	//fmt.Println(t.Field(3))
	//fmt.Println(t.Field(4))

	var missing []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		val, ok := m[field.Name]
		delete(m, field.Name)
		if !ok {
			missing = append(missing, field.Name)
			continue
		}

		switch field.Type.Kind() {
		// TODO: if the field is an integer you need to transform the val from float
		default:
			v.Field(i).Set(reflect.ValueOf(val))
		}
	}

	if len(missing) > 0 {
		return errors.New("missing fields: " + strings.Join(missing, ", "))
	}

	if len(m) > 0 {
		extra := make([]string, 0, len(m))
		for field := range m {
			extra = append(extra, field)
		}
		// TODO: consider sorting the output to get deterministic errors:
		// sort.Strings(extra)
		return errors.New("unknown fields: " + strings.Join(extra, ", "))
	}

	return nil
} */
