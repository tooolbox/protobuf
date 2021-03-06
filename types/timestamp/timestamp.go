package timestamp

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

// IsNull will return if the current timestamp is null
func (ts *Timestamp) IsNull() bool {

	if ts == nil {
		return true
	}

	// we use IsNotNull instead of IsNull to make sure that a timestamp is
	// initialized as null value
	return ts.IsNotNull == false
}

// Set will set the timestamp to the given time
func (ts *Timestamp) Set(value time.Time) {

	if ts == nil {
		*ts = Timestamp{}
	}

	if value.IsZero() {
		ts.SetNull()
		return
	}

	ts.Milliseconds = value.UnixNano() / 1000 / 1000
	ts.IsNotNull = true
}

// SetNull will clear the timestamp
func (ts *Timestamp) SetNull() {

	if ts == nil {
		return
	}

	ts.Milliseconds = 0
	ts.IsNotNull = false
}

// Time returns a golang time object
func (ts *Timestamp) Time() time.Time {

	if ts.IsNull() {
		return time.Time{}
	}

	return time.Unix(0, ts.Milliseconds*1000*1000)
}

// Scan implements the Scanner interface of the database driver
func (ts *Timestamp) Scan(value interface{}) error {

	// initialize timestamp if pointer is nil
	if ts == nil {
		*ts = Timestamp{}
	}

	// convert the interface to a time type
	dbTime, ok := value.(time.Time)

	if ok {
		ts.Milliseconds = dbTime.UnixNano() / 1000 / 1000
		ts.IsNotNull = true
		return nil
	}

	ts.Milliseconds = 0
	ts.IsNotNull = false
	return nil
}

// Value implements the db driver Valuer interface
func (ts Timestamp) Value() (driver.Value, error) {

	if ts.IsNull() {
		return nil, nil
	}

	return time.Unix(0, ts.Milliseconds*1000*1000), nil
}

// ImplementsGraphQLType is required by the graphql custom scalar interface
// this defines the name used in the schema to declare a null time type
func (ts *Timestamp) ImplementsGraphQLType(name string) bool {
	return name == "Time"
}

// UnmarshalGraphQL is required by the graphql custom scalar interface
// this wraps the null time
func (ts *Timestamp) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {

	case Timestamp:
		ts.IsNotNull = input.IsNotNull
		ts.Milliseconds = input.Milliseconds
		return nil

	case time.Time:
		time := &Timestamp{}
		time.Milliseconds = input.UnixNano() / 1000 / 1000
		time.IsNotNull = true
		return nil

	case string:

		timepoint, err := parseFromString(input)
		if err != nil {
			return err
		}

		ts.Set(timepoint)
		return nil

	case nil:
		ts.SetNull()
		return nil

	default:
		fmt.Printf("%T\n", input)
		fmt.Println(input)
		return fmt.Errorf("wrong type")
	}
}

// MarshalJSON will return the content as json value, this is also called
// by graphql to generate the response
func (ts Timestamp) MarshalJSON() ([]byte, error) {

	if ts.IsNull() {
		return []byte("null"), nil
	}

	// format the timestamp in iso compatible time format
	formatted := fmt.Sprintf("\"%s\"", ts.Time().Format(time.RFC3339))

	return []byte(formatted), nil
}

// UnmarshalJSON is used to convert the json representation into a timestamp
func (ts *Timestamp) UnmarshalJSON(input []byte) error {

	// trim the leading and trailing quotes from the timestamp
	cleanInput := bytes.Trim(input, "\"")

	// convert to string
	asString := string(cleanInput)

	if asString == "null" {
		ts.SetNull()
		return nil
	}

	timepoint, err := parseFromString(asString)
	if err != nil {
		return err
	}

	ts.Set(timepoint)
	return nil
}

// parseFromString will attemt to parse a timestamp string as time
func parseFromString(input string) (time.Time, error) {

	// try to parse the information as date
	timepoint, err := time.Parse(time.RFC3339, input)
	if err == nil {
		return timepoint, nil
	}

	// try to parse the information as date with nano precision from postgres
	timepoint, err = time.Parse("2006-01-02T15:04:05.999999", input)
	if err == nil {
		return timepoint, nil
	}

	// try to parse the information as unix timestamp
	asInt, err := strconv.ParseInt(input, 0, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("format for time must be RFC3339, RFC3339Nano format or unix timestamp: %s", input)
	}

	timepoint = time.Unix(asInt/1000, (asInt%1000)*1000*1000)
	return timepoint, nil

}
