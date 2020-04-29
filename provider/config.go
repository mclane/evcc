package provider

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/andig/evcc/util"
	"github.com/benbjohnson/clock"
	"github.com/savaki/jq"
)

// Config is the general provider config
type Config struct {
	Type  string
	Other map[string]interface{} `mapstructure:",remain"`
}

// MQTT singleton
var MQTT *MqttClient

func wrappedStringGetterFromConfig(log *util.Logger, stringG StringGetter, other map[string]interface{}) (
	StringGetter, map[string]interface{},
) {
	var cc struct {
		Jq    string
		Cache time.Duration
		Other map[string]interface{}
	}
	util.DecodeOther(log, other, &cc)

	// decorate cache
	if cc.Cache != 0 {
		clock := clock.New()
		var updated time.Time
		var val string

		stringG = StringGetter(func() (string, error) {
			if clock.Since(updated) > cc.Cache {
				new, err := stringG()
				if err != nil {
					return new, err
				}

				updated = clock.Now()
				val = new
			}

			return val, nil
		})
	}

	if cc.Jq == "" {
		return stringG, nil
	}

	jqOp, err := jq.Parse(cc.Jq)
	if err != nil {
		log.FATAL.Fatalf("config: invalid jq query: %s", cc.Jq)
	}

	// decorate jq
	return StringGetter(func() (string, error) {
		s, err := stringG()
		if err != nil {
			return s, err
		}

		b, err := jqOp.Apply([]byte(s))
		return string(b), err
	}), cc.Other
}

// stringGetterFromConfig creates a StringGetter from config
func stringGetterFromConfig(log *util.Logger, config Config) (
	res StringGetter, other map[string]interface{},
) {
	switch strings.ToLower(config.Type) {
	case "mqtt":
		if MQTT == nil {
			log.FATAL.Fatal("mqtt not configured")
		}

		var cc struct {
			Topic   string
			Timeout time.Duration
			Other   map[string]interface{}
		}
		util.DecodeOther(log, config.Other, &cc)
		other = cc.Other // remainder

		res = MQTT.StringGetter(cc.Topic, cc.Timeout)

	case "script":
		var cc struct {
			Cmd     string
			Timeout time.Duration
			Other   map[string]interface{}
		}
		util.DecodeOther(log, config.Other, &cc)
		other = cc.Other // remainder

		res = NewScriptProvider(log, cc.Cmd, cc.Timeout).StringGetter

	case "combined":
		res = openWBStatusFromConfig(log, config.Other)

	default:
		log.FATAL.Fatalf("invalid provider type %s", config.Type)
	}

	return wrappedStringGetterFromConfig(log, res, other)
}

// NewStringGetterFromConfig creates a StringGetter from config
func NewStringGetterFromConfig(log *util.Logger, config Config) StringGetter {
	stringG, other := stringGetterFromConfig(log, config)

	if len(other) > 0 {
		log.FATAL.Fatalf("config: unexpected config %+v", other)
	}

	return stringG
}

func floatGetterFromConfig(log *util.Logger, stringG StringGetter, other map[string]interface{}) (
	FloatGetter, map[string]interface{},
) {
	var cc struct {
		Scale float64
		Other map[string]interface{}
	}
	util.DecodeOther(log, other, &cc)

	return FloatGetter(func() (float64, error) {
		s, err := stringG()
		if err != nil {
			return 0, err
		}

		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, err
		}

		if cc.Scale > 0 {
			f *= cc.Scale
		}

		return f, nil
	}), cc.Other
}

// NewFloatGetterFromConfig creates a FloatGetter from config
func NewFloatGetterFromConfig(log *util.Logger, config Config) FloatGetter {
	stringG, other := stringGetterFromConfig(log, config)
	floatG, other := floatGetterFromConfig(log, stringG, other)

	if len(other) > 0 {
		log.FATAL.Fatalf("config: unexpected config %+v", other)
	}

	return floatG
}

// NewIntGetterFromConfig creates a IntGetter from config
func NewIntGetterFromConfig(log *util.Logger, config Config) IntGetter {
	floatG := NewFloatGetterFromConfig(log, config)

	return IntGetter(func() (int64, error) {
		f, err := floatG()
		return int64(math.Round(f)), err
	})
}

// NewBoolGetterFromConfig creates a BoolGetter from config
func NewBoolGetterFromConfig(log *util.Logger, config Config) BoolGetter {
	stringG := NewStringGetterFromConfig(log, config)

	return BoolGetter(func() (bool, error) {
		s, err := stringG()
		if err != nil {
			return false, err
		}
		return util.Truish(s), nil
	})
}

// NewIntSetterFromConfig creates a IntSetter from config
func NewIntSetterFromConfig(log *util.Logger, param string, config Config) (res IntSetter) {
	switch strings.ToLower(config.Type) {
	case "mqtt":
		if MQTT == nil {
			log.FATAL.Fatal("mqtt not configured")
		}

		var cc struct {
			Topic, Payload string
			Timeout        time.Duration
		}
		util.DecodeOther(log, config.Other, &cc)

		res = MQTT.IntSetter(param, cc.Topic, cc.Payload)

	case "script":
		var cc struct {
			Cmd     string
			Timeout time.Duration
		}
		util.DecodeOther(log, config.Other, &cc)

		exec := NewScriptProvider(log, cc.Cmd, cc.Timeout)
		res = exec.IntSetter(param)

	default:
		log.FATAL.Fatalf("invalid setter type %s", config.Type)
	}
	return
}

// NewBoolSetterFromConfig creates a BoolSetter from config
func NewBoolSetterFromConfig(log *util.Logger, param string, config Config) (res BoolSetter) {
	switch strings.ToLower(config.Type) {
	case "mqtt":
		if MQTT == nil {
			log.FATAL.Fatal("mqtt not configured")
		}

		var cc struct {
			Topic, Payload string
			Timeout        time.Duration
		}
		util.DecodeOther(log, config.Other, &cc)

		res = MQTT.BoolSetter(param, cc.Topic, cc.Payload)

	case "script":
		var cc struct {
			Cmd     string
			Timeout time.Duration
		}
		util.DecodeOther(log, config.Other, &cc)

		exec := NewScriptProvider(log, cc.Cmd, cc.Timeout)
		res = exec.BoolSetter(param)

	default:
		log.FATAL.Fatalf("invalid setter type %s", config.Type)
	}
	return
}
