// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *DefaultOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *DefaultOK) encodeFields(e *jx.Encoder) {
	{
		if s.Date.Set {
			e.FieldStart("date")
			s.Date.Encode(e, json.NewTimeEncoder("02/01/2006"))
		}
	}
	{
		if s.Time.Set {
			e.FieldStart("time")
			s.Time.Encode(e, json.NewTimeEncoder("3:04PM"))
		}
	}
	{
		if s.DateTime.Set {
			e.FieldStart("dateTime")
			s.DateTime.Encode(e, json.NewTimeEncoder("2006-01-02T15:04:05.999999999Z07:00"))
		}
	}
}

var jsonFieldsNameOfDefaultOK = [3]string{
	0: "date",
	1: "time",
	2: "dateTime",
}

// Decode decodes DefaultOK from json.
func (s *DefaultOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DefaultOK to nil")
	}
	s.setDefaults()

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "date":
			if err := func() error {
				s.Date.Reset()
				if err := s.Date.Decode(d, json.NewTimeDecoder("02/01/2006")); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"date\"")
			}
		case "time":
			if err := func() error {
				s.Time.Reset()
				if err := s.Time.Decode(d, json.NewTimeDecoder("3:04PM")); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"time\"")
			}
		case "dateTime":
			if err := func() error {
				s.DateTime.Reset()
				if err := s.DateTime.Decode(d, json.NewTimeDecoder("2006-01-02T15:04:05.999999999Z07:00")); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"dateTime\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode DefaultOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *DefaultOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DefaultOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes time.Time as json.
func (o OptDate) Encode(e *jx.Encoder, format func(*jx.Encoder, time.Time)) {
	if !o.Set {
		return
	}
	format(e, o.Value)
}

// Decode decodes time.Time from json.
func (o *OptDate) Decode(d *jx.Decoder, format func(*jx.Decoder) (time.Time, error)) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptDate to nil")
	}
	o.Set = true
	v, err := format(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptDate) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e, json.NewTimeEncoder("02/01/2006"))
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptDate) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d, json.NewTimeDecoder("02/01/2006"))
}

// Encode encodes time.Time as json.
func (o OptDateTime) Encode(e *jx.Encoder, format func(*jx.Encoder, time.Time)) {
	if !o.Set {
		return
	}
	format(e, o.Value)
}

// Decode decodes time.Time from json.
func (o *OptDateTime) Decode(d *jx.Decoder, format func(*jx.Decoder) (time.Time, error)) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptDateTime to nil")
	}
	o.Set = true
	v, err := format(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptDateTime) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e, json.NewTimeEncoder("2006-01-02T15:04:05.999999999Z07:00"))
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptDateTime) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d, json.NewTimeDecoder("2006-01-02T15:04:05.999999999Z07:00"))
}

// Encode encodes time.Time as json.
func (o OptTime) Encode(e *jx.Encoder, format func(*jx.Encoder, time.Time)) {
	if !o.Set {
		return
	}
	format(e, o.Value)
}

// Decode decodes time.Time from json.
func (o *OptTime) Decode(d *jx.Decoder, format func(*jx.Decoder) (time.Time, error)) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptTime to nil")
	}
	o.Set = true
	v, err := format(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptTime) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e, json.NewTimeEncoder("3:04PM"))
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptTime) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d, json.NewTimeDecoder("3:04PM"))
}

// Encode implements json.Marshaler.
func (s *RequiredOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *RequiredOK) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("date")
		json.EncodeTimeFormat(e, s.Date, "02/01/2006")
	}
	{
		e.FieldStart("time")
		json.EncodeTimeFormat(e, s.Time, "3:04PM")
	}
	{
		e.FieldStart("dateTime")
		json.EncodeTimeFormat(e, s.DateTime, "2006-01-02T15:04:05.999999999Z07:00")
	}
}

var jsonFieldsNameOfRequiredOK = [3]string{
	0: "date",
	1: "time",
	2: "dateTime",
}

// Decode decodes RequiredOK from json.
func (s *RequiredOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode RequiredOK to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "date":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := json.DecodeTimeFormat(d, "02/01/2006")
				s.Date = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"date\"")
			}
		case "time":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := json.DecodeTimeFormat(d, "3:04PM")
				s.Time = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"time\"")
			}
		case "dateTime":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := json.DecodeTimeFormat(d, "2006-01-02T15:04:05.999999999Z07:00")
				s.DateTime = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"dateTime\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode RequiredOK")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfRequiredOK) {
					name = jsonFieldsNameOfRequiredOK[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *RequiredOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *RequiredOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}