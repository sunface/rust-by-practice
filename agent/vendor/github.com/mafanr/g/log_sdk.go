//  Copyright © 2018 Sunface <CTO@188.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package g

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var L *zap.Logger
var DL *zap.Logger
var level string

func InitLogger(level string) {
	level = strings.ToLower(level)
	var lv zapcore.Level
	switch level {
	case "debug":
		lv = zap.DebugLevel
	case "info":
		lv = zap.InfoLevel
	case "warn":
		lv = zap.WarnLevel
	case "error":
		lv = zap.ErrorLevel
	}

	atom := zap.NewAtomicLevel()

	// To keep the example deterministic, disable timestamps in the output.
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	L = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),

		zapcore.Lock(os.Stdout),
		atom,
	), zap.AddCaller())

	atom.SetLevel(lv)

	atom1 := zap.NewAtomicLevel()

	// To keep the example deterministic, disable timestamps in the output.
	encoderCfg1 := zap.NewProductionEncoderConfig()
	encoderCfg1.EncodeTime = zapcore.ISO8601TimeEncoder
	DL = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg1),

		zapcore.Lock(os.Stdout),
		atom1,
	), zap.AddCaller())

	atom1.SetLevel(zap.DebugLevel)
}
