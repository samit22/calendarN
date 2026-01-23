package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConverEtoN(t *testing.T) {
	t.Log("when date is invalid it returns error")
	{
		date := "sdasd"

		_, err := converEtoN(date)
		expErr := `parsing time "sdasd" as "2006-01-02": cannot parse "sdasd" as "2006"`
		assert.Error(t, err, "expected error didnot receive")
		assert.Equal(t, err.Error(), expErr)
	}

	t.Log("when date is valid it returns converted date")
	{
		date := "2022-08-18"

		nDate, err := converEtoN(date)
		assert.NoError(t, err)
		assert.Equal(t, "2079-05-02", nDate)
	}
}

func Test_ConvertNtoE(t *testing.T) {
	t.Log("when date is invalid it returns error")
	{
		date := "sdasd"

		_, err := convertNtoE(date)
		expErr := `invalid date format, should be 2079-11-02`
		assert.Error(t, err, "expected error didnot receive")
		assert.Equal(t, expErr, err.Error())
	}

	t.Log("when date is valid it returns converted date")
	{
		date := "2079-05-02"

		eDate, err := convertNtoE(date)
		assert.NoError(t, err)
		assert.Equal(t, "2022-08-18", eDate)
	}

	t.Log("when nepali date 2067-05-02 is converted it returns 2010-08-18")
	{
		date := "2067-05-02"

		eDate, err := convertNtoE(date)
		assert.NoError(t, err)
		assert.Equal(t, "2010-08-18", eDate)
	}
}

func Test_DateConverter(t *testing.T) {
	t.Log("when argument does not have two values it returns error")
	{
		args := []string{}
		err := dateConvert(args)
		assert.Errorf(t, err, "expected error did not get one")
		assert.Equal(t, err.Error(), "argument does not include `etn` or `nte` and date")
	}

	t.Log("when arguemt has two values")
	{
		t.Log("when first argument is neither 'ent' nor 'nte' it returns error")
		{
			args := []string{"abc", "2022-08-18"}
			err := dateConvert(args)
			assert.Errorf(t, err, "expected error did not get one")
			assert.Equal(t, err.Error(), "argument is neither `etn` nor `nte`")
		}

		t.Log("when first argument is etn it converts english to nepali date")
		{
			args := []string{"etn", "2022-08-18"}
			err := dateConvert(args)
			assert.NoError(t, err, "unxpected error")
		}

		t.Log("when first argument is nte it converts nepali to english date")
		{
			args := []string{"nte", "2079-05-02"}
			err := dateConvert(args)
			assert.NoError(t, err, "unexpected error")
		}

	}
}
