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
	"net/http"
	cm "onlineAccountCreator/common"
	"time"
)

//CaptchaService service
type CaptchaService struct {
	Host string
}

//Captcha Captcha
type Captcha struct {
	Secret   string `json:"secret"`
	Response string `json:"response"`
	Remoteip string `json:"remoteip"`
}

//CaptchaResponse res
type CaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTs time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
	Code        int       `json:"code"`
}

//SendCaptchaCall SendCaptchaCall
func (c *CaptchaService) SendCaptchaCall(cap Captcha) *CaptchaResponse {
	var rtn = new(CaptchaResponse)
	var sURL = c.Host + "?secret=" + cap.Secret + "&response=" + cap.Response + "&remoteip=" + cap.Remoteip

	req, fail := cm.GetRequest(sURL, http.MethodPost, nil)
	if !fail {

		code := cm.ProcessServiceCall(req, &rtn)
		rtn.Code = code
	}
	return rtn
}
