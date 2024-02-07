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

package cmd

import (
	"4rc3us/gomodoro/enums"
	"fmt"
	"strings"
	"time"

	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var pomodoroTime string
var unit string

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set time for pomodoro",
	Long: `set time for pomodoro using
	Supported units: M, S, H

	Example: gomodoro set --time 25 --unit M`,
	Run: func(cmd *cobra.Command, args []string) {
		pomodoroTimeInt, err := strconv.Atoi(pomodoroTime)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		unit, err := enums.NewTimeUnit(strings.ToUpper(unit))
		if err != nil {
			fmt.Println(err, "\nvalid units: M, S, H, D")
			os.Exit(1)
		}

		fmt.Println("Pomodoro time set to:", pomodoroTimeInt, unit)
		timer := time.NewTimer(time.Duration(pomodoroTimeInt) * unit.ToTimeDuration())
		<-timer.C
		fmt.Println("Pomodoro finished")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringVarP(&pomodoroTime, "time", "t", "", "Time for pomodoro")
	setCmd.Flags().StringVarP(&unit, "unit", "u", "M", "Unit for time")

	err := setCmd.MarkFlagRequired("time")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
