// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: seatbid.go

package openrtb

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *SeatBid) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *SeatBid) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "bid":`)
	if j.Bid != nil {
		buf.WriteString(`[`)
		for i, v := range j.Bid {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				err = v.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteByte(',')
	if len(j.Seat) != 0 {
		buf.WriteString(`"seat":`)
		fflib.WriteJsonString(buf, string(j.Seat))
		buf.WriteByte(',')
	}
	if j.Group != 0 {
		buf.WriteString(`"group":`)
		fflib.FormatBits2(buf, uint64(j.Group), 10, j.Group < 0)
		buf.WriteByte(',')
	}
	if len(j.Ext) != 0 {
		buf.WriteString(`"ext":`)

		{

			obj, err = j.Ext.MarshalJSON()
			if err != nil {
				return err
			}
			buf.Write(obj)

		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtSeatBidbase = iota
	ffjtSeatBidnosuchkey

	ffjtSeatBidBid

	ffjtSeatBidSeat

	ffjtSeatBidGroup

	ffjtSeatBidExt
)

var ffjKeySeatBidBid = []byte("bid")

var ffjKeySeatBidSeat = []byte("seat")

var ffjKeySeatBidGroup = []byte("group")

var ffjKeySeatBidExt = []byte("ext")

// UnmarshalJSON umarshall json - template of ffjson
func (j *SeatBid) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *SeatBid) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtSeatBidbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtSeatBidnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'b':

					if bytes.Equal(ffjKeySeatBidBid, kn) {
						currentKey = ffjtSeatBidBid
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeySeatBidExt, kn) {
						currentKey = ffjtSeatBidExt
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'g':

					if bytes.Equal(ffjKeySeatBidGroup, kn) {
						currentKey = ffjtSeatBidGroup
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeySeatBidSeat, kn) {
						currentKey = ffjtSeatBidSeat
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeySeatBidExt, kn) {
					currentKey = ffjtSeatBidExt
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeySeatBidGroup, kn) {
					currentKey = ffjtSeatBidGroup
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeySeatBidSeat, kn) {
					currentKey = ffjtSeatBidSeat
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeySeatBidBid, kn) {
					currentKey = ffjtSeatBidBid
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtSeatBidnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtSeatBidBid:
					goto handle_Bid

				case ffjtSeatBidSeat:
					goto handle_Seat

				case ffjtSeatBidGroup:
					goto handle_Group

				case ffjtSeatBidExt:
					goto handle_Ext

				case ffjtSeatBidnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Bid:

	/* handler: j.Bid type=[]openrtb.Bid kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Bid = nil
		} else {

			j.Bid = []Bid{}

			wantVal := true

			for {

				var tmpJBid Bid

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJBid type=openrtb.Bid kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						err = tmpJBid.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
						if err != nil {
							return err
						}
					}
					state = fflib.FFParse_after_value
				}

				j.Bid = append(j.Bid, tmpJBid)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Seat:

	/* handler: j.Seat type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Seat = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Group:

	/* handler: j.Group type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.Group = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Ext:

	/* handler: j.Ext type=openrtb.Extension kind=slice quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Ext.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
