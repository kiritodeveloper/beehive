/*
 *    Copyright (C) 2014 Stefan 'glaxx' Luecke     
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Stefan 'glaxx' Luecke <glaxx@glaxx.net>            
 */

package timebee

import (
	"github.com/muesli/beehive/modules"
)

type TimeBeeFactory struct {
	modules.ModuleFactory
}

func (factory *TimeBeeFactory) New(name, description string, options modules.BeeOptions) modules.ModuleInterface {
	bee := TimeBee{
		second: options.GetValue("Second").(string),
		minute: options.GetValue("Minute").(string),
		dayofweek: options.GetValue("DayOfWeek").(string),
		dayofmonth: options.GetValue("DayOfMonth").(string),
		month: options.GetValue("Month").(string),
		year: options.GetValue("Year").(string),
	}
	bee.Module = modules.Module{name, factory.Name(), description}
	return &bee
}

func (factory *TimeBeeFactory) Name() string {
	return "timebee"
}

func (factory *TimeBeeFactory) Description() string {
	return "A bee that triggers an event at a given time"
}

func (factory *TimeBeeFactory) Options() []modules.BeeOptionDescriptor {
	opts := []modules.BeeOptionDescriptor{
		modules.BeeOptionDescriptor{
			Name:		"Second",
			Description:	"00-59 for a specific time or * for ignore",
			Type:		"string",
		},
		modules.BeeOptionDescriptor{
			Name:		"Minute",
			Description:	"00-59 for a specific time or * for ignore",
			Type:		"string",
		},
		modules.BeeOptionDescriptor{
			Name:		"DayOfWeek",
			Descripton:	"0-6 0 = Monday 6 = Sunday, * for ignore";
			Type:		"string"
		},
		modules.BeeOptionDescriptor{
			Name:		"DayOfMonth",
			Description:	"01-31 for a specific time or * for ignore)",
			Type:		"string",
		},
		modules.BeeOptionDescriptor{
			Name:		"Month",
			Description:	"01 - 12 for a specific time or * for ignore)",
			Type:		"string",
		},
		modules.BeeOptionDescriptor{
			Name:		"Year",
			Description:	"2014 - 9999 for specific time (non-reoccuring) or * for ignore (recommended)",
			Type:		"string",
		},
	}
	return opts
}

func (factory *TimeBeeFactory) Events() []modules.EventDescriptor {
	events := []modules.EventDescriptor{
		modules.EventDescriptor{
			Namespace:	factory.Name(),
			Name:		"time_event",
			Description:	"The time has come ...",
			Options: []modules.PlaceholderDescriptor{},
			},
		}
	return events
}
/*
func (factory *TimeBeeFactory) Actions() []modules.ActionDescriptor {
        actions := []modules.ActionDescriptor{}
        return actions
}

func (factory *TimeBeeFactory) Image() string {
	return ""
}*/

func init() {
	f := TimeBeeFactory{}
	modules.RegisterFactory(&f)
}
