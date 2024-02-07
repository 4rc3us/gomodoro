/*
Copyright © 2024 J Daniel Arce <juandanielarce398@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package enums

import "errors"

type Unit string

const (
	Minutes Unit = "M"
	Hours   Unit = "H"
	Seconds Unit = "S"
	Days    Unit = "D"
)

func NewTimeUnit(newUnit string) (Unit, error) {
	unites := []string{"M", "H", "S", "D"}

	for _, v := range unites {
		if v == newUnit {
			return Unit(newUnit), nil
		}
	}

	err := errors.New("Invalid time unit")
	return "", err
}
