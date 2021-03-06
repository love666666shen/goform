package goform

type MultiCheckboxElement struct {
	Element
}

func NewMultiCheckboxElement(name string, label string, attributes []*Attribute, valueOptions []*ValueOption, validators []ValidatorInterface, filters []FilterInterface) *MultiCheckboxElement {
	element := new(MultiCheckboxElement)
	element.Type = ElementTypeMultiCheckbox
	element.Name = name
	element.Label = label
	element.Attributes = attributes
	element.ValueOptions = valueOptions
	element.Validators = validators
	element.Filters = filters

	return element
}

func (element *MultiCheckboxElement) Render() string {
	return renderTemplate(ElementTypeMultiCheckbox, element)
}
