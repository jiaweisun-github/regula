package tests.rules.REGULA_R00003_infra

import data.fugue.regula

test_REGULA_R00003 {
  report := regula.report with input as mock_input
  resources := report.rules.REGULA_R00003.resources

  resources["aws_vpc.valid_vpc"].valid == true
  resources["aws_vpc.invalid_vpc"].valid == false
}