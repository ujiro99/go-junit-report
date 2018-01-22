package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/ujiro99/doctest-junit-report/parser"
)

type TestCase struct {
	name        string
	reportName  string
	report      *parser.Report
	noXMLHeader bool
	packageName string
}

var testCases = []TestCase{
	{
		name:       "alternative_macros.txt",
		reportName: "alternative_macros.xml",
		report: &parser.Report{
			TestSuites: []*parser.TestSuite{
				{
					Name: "alternative_macros.cpp",
					Time: 414,
					Tests: []*parser.Test{
						{
							Name:     "custom macros",
							Time:     414,
							Result:   parser.PASS,
							Output:   []string{},
							Filename: "alternative_macros.cpp",
						},
					},
				},
			},
		},
	}, {
		name:       "assertion_macros.txt",
		reportName: "assertion_macros.xml",
		report: &parser.Report{
			TestSuites: []*parser.TestSuite{
				{
					Name: "assertion_macros.cpp",
					Time: 4122,
					Tests: []*parser.Test{
						{
							Name:     "normal macros",
							Time:     993,
							Result:   parser.FAIL,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "expressions should be evaluated only once",
							Time:     61,
							Result:   parser.PASS,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "exceptions-related macros",
							Time:     224,
							Result:   parser.FAIL,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "exceptions-related macros for std::exception",
							Time:     226,
							Result:   parser.FAIL,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "WARN level of asserts don't fail the test case",
							Time:     481,
							Result:   parser.PASS,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "CHECK level of asserts fail the test case but don't abort it",
							Time:     248,
							Result:   parser.FAIL,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "REQUIRE level of asserts fail and abort the test case - 1",
							Time:     115,
							Result:   parser.FAIL,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "REQUIRE level of asserts fail and abort the test case - 2",
							Time:     53,
							Result:   parser.FAIL,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "REQUIRE level of asserts fail and abort the test case - 3",
							Time:     48,
							Result:   parser.FAIL,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "REQUIRE level of asserts fail and abort the test case - 4",
							Time:     50,
							Result:   parser.FAIL,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "REQUIRE level of asserts fail and abort the test case - 5",
							Time:     86,
							Result:   parser.FAIL,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "REQUIRE level of asserts fail and abort the test case - 6",
							Time:     61,
							Result:   parser.FAIL,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "REQUIRE level of asserts fail and abort the test case - 7",
							Time:     50,
							Result:   parser.FAIL,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "REQUIRE level of asserts fail and abort the test case - 8",
							Time:     55,
							Result:   parser.FAIL,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "REQUIRE level of asserts fail and abort the test case - 9",
							Time:     50,
							Result:   parser.FAIL,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "REQUIRE level of asserts fail and abort the test case - 10",
							Time:     50,
							Result:   parser.FAIL,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "REQUIRE level of asserts fail and abort the test case - 11",
							Time:     47,
							Result:   parser.FAIL,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "REQUIRE level of asserts fail and abort the test case - 12",
							Time:     48,
							Result:   parser.FAIL,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "all binary assertions",
							Time:     904,
							Result:   parser.PASS,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
						{
							Name:     "some asserts used in a function called by a test case",
							Time:     272,
							Result:   parser.PASS,
							Output:   []string{},
							Filename: "assertion_macros.cpp",
						},
					},
				},
			},
		},
	},
}

func TestParser(t *testing.T) {
	for _, testCase := range testCases {
		t.Logf("Running: %s", testCase.name)

		file, err := os.Open("tests/" + testCase.name)
		if err != nil {
			t.Fatal(err)
		}

		report, err := parser.Parse(file)
		if err != nil {
			t.Fatalf("error parsing: %s", err)
		}

		if report == nil {
			t.Fatalf("Report == nil")
		}

		expected := testCase.report
		if len(report.TestSuites) != len(expected.TestSuites) {
			t.Fatalf("Report packages == %d, want %d", len(report.TestSuites), len(expected.TestSuites))
		}

		for i, pkg := range report.TestSuites {
			expPkg := expected.TestSuites[i]

			if pkg.Name != expPkg.Name {
				t.Logf("TestSuite.Name %s", pkg.Name)
				t.Errorf("TestSuite.Name == %s, want %s", pkg.Name, expPkg.Name)
			}

			if pkg.Time != expPkg.Time {
				t.Logf("TestSuite.Name %s", pkg.Name)
				t.Errorf("TestSuite.Time == %d, want %d", pkg.Time, expPkg.Time)
			}

			if len(pkg.Tests) != len(expPkg.Tests) {
				t.Logf("TestSuite.Name %s", pkg.Name)
				t.Fatalf("TestSuite Tests == %d, want %d", len(pkg.Tests), len(expPkg.Tests))
			}

			for j, test := range pkg.Tests {
				expTest := expPkg.Tests[j]

				if test.Name != expTest.Name {
					t.Logf("Test.Name %s", test.Name)
					t.Errorf("Test.Name == %s, want %s", test.Name, expTest.Name)
				}

				if test.Time != expTest.Time {
					t.Logf("Test.Name %s", test.Name)
					t.Errorf("Test.Time == %d, want %d", test.Time, expTest.Time)
				}

				if test.Result != expTest.Result {
					t.Logf("Test.Name %s", test.Name)
					t.Errorf("Test.Result == %d, want %d", test.Result, expTest.Result)
				}

				testOutput := strings.Join(test.Output, "\n")
				expTestOutput := strings.Join(expTest.Output, "\n")
				if testOutput != expTestOutput {
					t.Logf("Test.Name %s", test.Name)
					t.Errorf("Test.Output (%s) ==\n%s\n, want\n%s", test.Name, testOutput, expTestOutput)
				}
			}
		}
	}
}

func TestJUnitFormatter(t *testing.T) {
	testJUnitFormatter(t, "", "")
}

func TestVersionFlag(t *testing.T) {
	testJUnitFormatter(t, "custom-version", "")
}

func TestPackageNameFlag(t *testing.T) {
	testJUnitFormatter(t, "", "custom-package")
}

func testJUnitFormatter(t *testing.T, version string, packageName string) {
	for _, testCase := range testCases {
		report, err := loadTestReport(testCase.reportName, version, packageName)
		if err != nil {
			t.Fatal(err)
		}

		var junitReport bytes.Buffer

		if err = JUnitReportXML(testCase.report, testCase.noXMLHeader, version, packageName, &junitReport); err != nil {
			t.Fatal(err)
		}

		if string(junitReport.Bytes()) != report {
			t.Fatalf("Fail: %s Report xml ==\n%s, want\n%s", testCase.name, string(junitReport.Bytes()), report)
		}
	}
}

func loadTestReport(name, version string, packageName string) (string, error) {
	contents, err := ioutil.ReadFile("tests/" + name)
	if err != nil {
		return "", err
	}

	report := string(contents)

	// replace value="1.0" With specified version
	if version != "" {
		report = strings.Replace(report, `<properties></properties>`, fmt.Sprintf(`<properties>
		<property name="version" value="%v"></property>
	</properties>`, version), -1)
	}

	// replace value="package" With specified packageName
	if packageName != "" {
		report = strings.Replace(report, `">`, fmt.Sprintf(`" name="%s">`, packageName), 1)
	}

	return report, nil
}
