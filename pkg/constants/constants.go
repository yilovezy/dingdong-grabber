/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at
  http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package constants

const (
	CookiePrefix = "DDXQSESSID"
	SignNars     = "nars"
	SignSesi     = "sesi"
	Sign         = "sign"
	ImSecret     = "im_secret"

	Secret   = "Y29uc3QgeyBhcmd2IH0gPSByZXF1aXJlKCdwcm9jZXNzJyk7Cgp2YXIgdG9vbHMgPSB7CiAgICBwYXJzZV9xdWVyeTogKHF1ZXJ5KSA9PiB7CiAgICAgICAgdmFyIHF1ZXJ5X29iaiA9IHt9OwogICAgICAgIGlmIChxdWVyeS5pbmRleE9mKCc/JykgIT09IC0xKSBxdWVyeSA9IHF1ZXJ5LnNwbGl0KCc/JylbMV07CiAgICAgICAgcXVlcnkuc3BsaXQoJyYnKS5mb3JFYWNoKChpdGVtKSA9PiB7CiAgICAgICAgICAgIHZhciBba2V5LCB2YWx1ZV0gPSBpdGVtLnNwbGl0KCc9Jyk7CiAgICAgICAgICAgIHF1ZXJ5X29ialtrZXldID0gdmFsdWU7CiAgICAgICAgfSk7CiAgICAgICAgcmV0dXJuIHF1ZXJ5X29iajsKICAgIH0sCiAgICByYW5kb21fc3RyaW5nOiAobGVuKSA9PiB7CiAgICAgICAgdmFyIHN0ciA9ICdhYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5ekFCQ0RFRkdISUpLTE1OT1BRUlNUVVZXWFlaMDEyMzQ1Njc4OSc7CiAgICAgICAgdmFyIHJlc3VsdCA9ICcnOwogICAgICAgIGZvciAodmFyIGkgPSAwOyBpIDwgbGVuOyBpKyspIHsKICAgICAgICAgICAgcmVzdWx0ICs9IHN0cltNYXRoLmZsb29yKE1hdGgucmFuZG9tKCkgKiBzdHIubGVuZ3RoKV07CiAgICAgICAgfQogICAgICAgIHJldHVybiByZXN1bHQ7CiAgICB9LAogICAgcmFuZF9iZXR3ZWVuOiAobWluLCBtYXgpID0+IHsKICAgICAgICByZXR1cm4gTWF0aC5mbG9vcihNYXRoLnJhbmRvbSgpICogKG1heCAtIG1pbiArIDEpICsgbWluKTsKICAgIH0sCiAgICBzbGVlcDogKG1zKSA9PiB7CiAgICAgICAgcmV0dXJuIG5ldyBQcm9taXNlKChyZXNvbHZlKSA9PiBzZXRUaW1lb3V0KHJlc29sdmUsIG1zKSk7CiAgICB9LAp9OwoKdmFyIHNpZ25fZnVuYyA9ICh0LCBlKSA9PiB7CiAgICB2YXIgUyA9IHJlcXVpcmUoJ2NyeXB0bycpOwogICAgZSA9IEpTT04ucGFyc2UoSlNPTi5zdHJpbmdpZnkoZSkpOwogICAgZVsncHJpdmF0ZV9rZXknXSA9IHQ7CiAgICB2YXIgcyA9IE9iamVjdC5rZXlzKGUpCiAgICAgICAgLnNvcnQoKQogICAgICAgIC5tYXAoKHQpID0+IHsKICAgICAgICAgICAgcmV0dXJuIHQgKyAnPScgKyBlW3RdOwogICAgICAgIH0pCiAgICAgICAgLmpvaW4oJyYnKTsKICAgIHZhciByID0gUy5jcmVhdGVIYXNoKCdtZDUnKQogICAgICAgIC51cGRhdGUocykKICAgICAgICAuZGlnZXN0KCkKICAgICAgICAucmVkdWNlKCh0LCBlKSA9PiB7CiAgICAgICAgICAgIHJldHVybiAoCiAgICAgICAgICAgICAgICB0ICsKICAgICAgICAgICAgICAgICgoZSAmIDI1NSkudG9TdHJpbmcoMTYpLmxlbmd0aCA9PT0gMQogICAgICAgICAgICAgICAgICAgID8gJzAnICsgKGUgJiAyNTUpLnRvU3RyaW5nKDE2KQogICAgICAgICAgICAgICAgICAgIDogKGUgJiAyNTUpLnRvU3RyaW5nKDE2KSkKICAgICAgICAgICAgKTsKICAgICAgICB9LCAnJykKICAgICAgICAudG9Mb3dlckNhc2UoKTsKICAgIHZhciBhID0gKHQpID0+IHsKICAgICAgICB2YXIgZSA9IFMuY3JlYXRlSGFzaCgnbWQ1JykudXBkYXRlKHQpLmRpZ2VzdCgnaGV4JykudG9Mb3dlckNhc2UoKTsKICAgICAgICB2YXIgcyA9IGUuc3Vic3RyKDMsIDkpOwogICAgICAgIHZhciByID0gZS5zdWJzdHIoNywgMjApOwogICAgICAgIHZhciBhID0gZS5zdWJzdHIoMSwgOSk7CiAgICAgICAgdmFyIG8gPSB0CiAgICAgICAgICAgIC5zcGxpdCgnJykKICAgICAgICAgICAgLm1hcCgodCkgPT4gewogICAgICAgICAgICAgICAgcmV0dXJuIHQuY2hhckNvZGVBdCgwKS50b1N0cmluZygxNik7CiAgICAgICAgICAgIH0pCiAgICAgICAgICAgIC5qb2luKCcnKTsKICAgICAgICB2YXIgbCA9IHRvb2xzLnJhbmRvbV9zdHJpbmcoMyk7CiAgICAgICAgdmFyIG4gPSBbcywgciwgYSwgb10uc29ydCgpLmpvaW4oJycpICsgbDsKICAgICAgICB2YXIgaSA9IFMuY3JlYXRlSGFzaCgnc2hhMScpLnVwZGF0ZShuKS5kaWdlc3QoJ2hleCcpLnRvTG93ZXJDYXNlKCk7CiAgICAgICAgdmFyIHUgPSBpICsgbDsKICAgICAgICB2YXIgZCA9IG87CiAgICAgICAgdmFyIGggPSBlLnN1YnN0cigzLCAyNik7CiAgICAgICAgdmFyIGcgPSBkICsgaDsKICAgICAgICB2YXIgcCA9IGcKICAgICAgICAgICAgLnNwbGl0KCcnKQogICAgICAgICAgICAubWFwKCh0KSA9PiB7CiAgICAgICAgICAgICAgICByZXR1cm4gdC5jaGFyQ29kZUF0KDApLnRvU3RyaW5nKDE2KTsKICAgICAgICAgICAgfSkKICAgICAgICAgICAgLmpvaW4oJycpCiAgICAgICAgICAgIC5zdWJzdHIoMSwgMTcpOwogICAgICAgIHZhciBiID0gQnVmZmVyLmZyb20odSArIHApLnRvU3RyaW5nKCdiYXNlNjQnKTsKICAgICAgICB2YXIgYyA9CiAgICAgICAgICAgIGIgKwogICAgICAgICAgICBnCiAgICAgICAgICAgICAgICAuc3BsaXQoJycpCiAgICAgICAgICAgICAgICAubWFwKCh0KSA9PiB7CiAgICAgICAgICAgICAgICAgICAgcmV0dXJuIHQuY2hhckNvZGVBdCgwKS50b1N0cmluZygxNik7CiAgICAgICAgICAgICAgICB9KQogICAgICAgICAgICAgICAgLmpvaW4oJycpCiAgICAgICAgICAgICAgICAuc3Vic3RyKDQ4LCA2NCkgKwogICAgICAgICAgICAnaGV4YmFzZTY0JzsKICAgICAgICB2YXIgbSA9IFMuY3JlYXRlSGFzaCgnbWQ1JykudXBkYXRlKGMpLmRpZ2VzdCgnaGV4JykudG9Mb3dlckNhc2UoKTsKICAgICAgICByZXR1cm4gbCArIGggKyBtOwogICAgfTsKICAgIHZhciBvID0gKHQpID0+IHsKICAgICAgICB2YXIgZSA9IFMuY3JlYXRlSGFzaCgnbWQ1JykudXBkYXRlKHQpLmRpZ2VzdCgnaGV4JykudG9Mb3dlckNhc2UoKTsKICAgICAgICB2YXIgcyA9IGUuc3Vic3RyKDUsIDIxKTsKICAgICAgICB2YXIgciA9IGU7CiAgICAgICAgdmFyIGEgPSByCiAgICAgICAgICAgIC5zcGxpdCgnJykKICAgICAgICAgICAgLm1hcCgodCkgPT4gewogICAgICAgICAgICAgICAgcmV0dXJuIHQuY2hhckNvZGVBdCgwKS50b1N0cmluZygxNik7CiAgICAgICAgICAgIH0pCiAgICAgICAgICAgIC5qb2luKCcnKQogICAgICAgICAgICAuc3Vic3RyKDEsIDQ0KTsKICAgICAgICB2YXIgbyA9IFMuY3JlYXRlSGFzaCgnbWQ1JykudXBkYXRlKGEpLmRpZ2VzdCgnaGV4JykudG9Mb3dlckNhc2UoKTsKICAgICAgICB2YXIgbCA9IG8uc3Vic3RyKDUsIDE2KTsKICAgICAgICB2YXIgbiA9IG8uc3Vic3RyKDAsIDE2KTsKICAgICAgICB2YXIgaSA9IHMgKyByICsgdDsKICAgICAgICB2YXIgdSA9IFMuY3JlYXRlQ2lwaGVyaXYoJ2Flcy0xMjgtY2JjJywgbCwgbik7CiAgICAgICAgdmFyIGQgPSB1LnVwZGF0ZShpLCAndXRmOCcsICdiYXNlNjQnKTsKICAgICAgICB2YXIgaCA9IHUuZmluYWwoJ2Jhc2U2NCcpOwogICAgICAgIHZhciBnID0gZCArIGg7CiAgICAgICAgdmFyIHAgPSByLnN1YnN0cigzLCA5KTsKICAgICAgICB2YXIgYiA9IHIuc3Vic3RyKDcsIDIwKTsKICAgICAgICB2YXIgYyA9IHIuc3Vic3RyKDgsIDkpOwogICAgICAgIHZhciBtID0KICAgICAgICAgICAgUy5jcmVhdGVIYXNoKCdzaGExJykKICAgICAgICAgICAgICAgIC51cGRhdGUoW3AsIGIsIGMsIGddLnNvcnQoKS5qb2luKCcnKSkKICAgICAgICAgICAgICAgIC5kaWdlc3QoJ2hleCcpCiAgICAgICAgICAgICAgICAudG9Mb3dlckNhc2UoKSArIHRvb2xzLnJhbmRvbV9zdHJpbmcoNik7CiAgICAgICAgdmFyIEMgPSB0b29scy5yYW5kb21fc3RyaW5nKDUpOwogICAgICAgIHJldHVybiBDICsgbTsKICAgIH07CiAgICByZXR1cm4geyBzaWduOiByLCBzZXNpOiBvKHIpLCBuYXJzOiBhKHIpIH07Cn07Cgpjb25zdCByZXN1bHQgPSBzaWduX2Z1bmMocHJvY2Vzcy5hcmd2WzJdLCBKU09OLnBhcnNlKHByb2Nlc3MuYXJndlszXSkpOwoKY29uc29sZS5sb2coSlNPTi5zdHJpbmdpZnkocmVzdWx0KSk7"
	SignFile = "sign.js"
)
