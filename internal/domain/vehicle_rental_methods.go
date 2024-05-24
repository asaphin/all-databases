package domain

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

func (a Address) String() string {
	var builder strings.Builder

	var err error

	if a.InCareOfName != "" {
		_, err = fmt.Fprintf(&builder, "In care of name %s, ", a.InCareOfName)
		if err != nil {
			log.WithError(err).Error("unable to add name info to address string")
		}
	}

	_, err = fmt.Fprintf(&builder, "%s, %s", a.StreetNumber, a.Street)
	if err != nil {
		log.WithError(err).Error("unable to add street info to address string")
	}
	if a.Apartment != "" {
		_, err = fmt.Fprintf(&builder, ", Apt %s", a.Apartment)
		if err != nil {
			log.WithError(err).Error("unable to add apartment info to address string")
		}
	}

	if len(a.AdditionalInfo) > 0 {
		builder.WriteString(" (")
		fields := make([]string, 0, len(a.AdditionalInfo))
		for key, value := range a.AdditionalInfo {
			fields = append(fields, fmt.Sprintf("%s: %s", key, value))
		}
		builder.WriteString(strings.Join(fields, ", "))
		builder.WriteString(")")
	}

	_, err = fmt.Fprintf(&builder, ", %s, %s, %s, %s", a.Locality, a.Region, a.Country, a.PostalCode)
	if err != nil {
		log.WithError(err).Error("unable to add locality info to address string")
	}

	_, err = fmt.Fprintf(&builder, " Lat: %.6f, Long: %.6f", a.Latitude, a.Longitude)
	if err != nil {
		log.WithError(err).Error("unable to add coordinates info to address string")
	}

	builder.WriteString(fmt.Sprintf(" (%s)", a.Type))

	return builder.String()
}

func (v Vehicle) String() string {
	return fmt.Sprintf("%s: %s %s, Model year: %d, S/N: %s %s", v.Type, v.Manufacturer, v.Model, v.Year, v.SerialNumber, mapToString(v.Metadata))
}

func mapToString(m map[string]interface{}) string {
	var builder strings.Builder
	builder.WriteString("{")
	first := true
	for key, value := range m {
		if !first {
			builder.WriteString(", ")
		} else {
			first = false
		}
		builder.WriteString(key)
		builder.WriteString(": ")
		switch v := value.(type) {
		case string:
			builder.WriteString(fmt.Sprintf(`"%s"`, v))
		case map[string]interface{}:
			builder.WriteString(mapToString(v))
		default:
			b, _ := json.Marshal(value)
			builder.Write(b)
		}
	}
	builder.WriteString("}")
	return builder.String()
}
