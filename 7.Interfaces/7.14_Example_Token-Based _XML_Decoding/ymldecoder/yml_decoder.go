package yml

import (
	"encoding/xml"
	"errors"
	"io"

	// "offers/internal/entity"
	// "offers/internal/usecase/decoder"
	"regexp"
	"strings"
)

func (y *ReaderYml) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	defer close(y.elements)
	for {

		select {
		case <-y.serverCtx.Done():
			break
		default:
		}

		t, err := d.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			} else {
				return err
			}
		}
		switch token := t.(type) {
		case xml.StartElement:
			if token.Name.Local == categoryElement {
				//просто на структуру
				err = d.DecodeElement(&y.category, &token)
				if err != nil {
					y.elements <- err
					continue
				}
				y.elements <- y.category
				y.category = entity.Category{}
				continue
			}
			//else
			err = y.startToken(token)
			if err != nil {
				y.elements <- err
				continue
			}
		case xml.CharData:
			err = y.datToken(token)
			if err != nil {
				y.elements <- err
				continue
			}
			continue
		case xml.EndElement:
			//Закрытие тегов
			err = y.endToken(token)
			if err != nil {
				y.elements <- err
				continue
			}
		default:
			continue
		}
	}
}

func (y *ReaderYml) datToken(t xml.CharData) (err error) {
	//data внутри элемента
	if !y.isOffers && !y.isParams {
		return nil
	}
	var matchedUnicode bool
	matchedUnicode, err = regexp.Match(patternRegExpUnicode, []byte(y.paramName))
	if err != nil {
		return err
	}
	if !matchedUnicode {
		return nil
	}
	value := strings.Trim(string(t), "\n\t\r")
	if value == "" {
		return nil
	}
	if y.uintParamValue != "" {
		value += " " + y.uintParamValue
	}
	if v, ok := y.fields[y.paramName]; ok {
		y.fields[y.paramName] = strings.Trim(v+","+value, ",")
	} else {
		y.fields[y.paramName] = value
	}
	y.paramName = ""
	return nil

}

func (y *ReaderYml) endToken(t xml.EndElement) error {
	//Закрытие тегов
	if t.Name.Local == offersElement {
		//Завершили обход всех товаров, тег </offer>
		y.isOffers = false
	}
	if t.Name.Local == paramElement {
		//Завершили обход параметров тег </param>
		y.isParams = false
		y.paramName = ""
		y.uintParamValue = ""
	}
	if t.Name.Local == offerElement {
		//Завершили обход товара, тег </offer>
		y.paramName = ""
		y.uintParamValue = ""
		offer, err := decoder.OfferDecodeFromMap(y.fields)
		if err != nil {
			return err
		}
		offer.Params = y.fields
		y.elements <- offer
		y.fields = make(map[string]string)
	}
	return nil
}

func (y *ReaderYml) startToken(t xml.StartElement) (err error) {
	//Открытие тега
	switch t.Name.Local {
	case offerElement:
		//начинаем обход товара
		for _, attr := range t.Attr {
			value := strings.Trim(attr.Value, "\n\t\r")
			if v, ok := y.fields[attr.Name.Local]; ok {
				y.fields[attr.Name.Local] = strings.Trim(v+","+value, ",")
				return nil
			}
			y.fields[attr.Name.Local] = value

		}
		return nil
	case paramElement:
		y.isParams = true
		for _, attr := range t.Attr {
			value := strings.Trim(attr.Value, "\n\t\r")
			if attr.Name.Local == unitElement {
				y.uintParamValue = value
				continue
			}
			y.paramName = value
		}
		return nil
	case offersElement:
		//Открылся тег <offers>
		y.isOffers = true
		return nil
	default:
		if y.isOffers {
			//если мы обходим товары
			//забираем название переменной
			var matchedUnicode bool
			matchedUnicode, err = regexp.Match(patternRegExpUnicode, []byte(t.Name.Local))
			if err != nil {
				return err
			}
			if !matchedUnicode {
				return nil
			}
			y.paramName = strings.Trim(t.Name.Local, "\n\t\n")
		}
		return nil
	}
}
