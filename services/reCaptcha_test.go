/*
 Copyright (C) 2018 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2018 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs LLC., or third
 parties.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published
 by the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package services

import (
	"fmt"
	"testing"
)

func TestCaptchaService_SendCaptchaCall(t *testing.T) {
	var c CaptchaService
	c.Host = "https://www.google.com/recaptcha/api/siteverify"
	var cp Captcha
	cp.Secret = "abaap"
	cp.Remoteip = "10.0.0.1"
	cp.Response = "lsljdiididi"

	res := c.SendCaptchaCall(cp)
	fmt.Print("Google ReCaptcha Resp: ")
	fmt.Println(res)
	if res.Success == true {
		t.Fail()
	}
}
