package envci

import (
	"strings"

	env "github.com/tj/go/env"
)

// Environment Variables
type Environment struct {
	Base   string
	Build  string
	Branch string
	Commit string
	Job    string
	PR     string
	Root   string
	Slug   string
}

// Services List
var Services = map[string]*Environment{
	"appveyor":  appveyor,
	"bamboo":    bamboo,
	"buildkite": buildKite,
	"circleci":  circleCI,
	"codeship":  codeship,
	"drone":     drone,
	"gitlab":    gitlab,
	"jenkins":   jenkins,
	"semaphore": semaphore,
	"shippable": shippable,
	"teamcity":  teamcity,
	"travis":    travis,
	"wercker":   wercker,
}

// pullRequestIsFalse returns "" if env equals "false"
func pullRequestIsFalse(name string) string {
	pr := env.GetDefault(name, "")
	if pr == "false" {
		return ""
	}

	return pr
}

// Appveyor Service - https://www.appveyor.com/docs/environment-variables/
var appveyor = &Environment{
	Base:   "APPVEYOR",
	Build:  env.GetDefault("APPVEYOR_BUILD_NUMBER", ""),
	Branch: env.GetDefault("APPVEYOR_REPO_BRANCH", ""),
	Commit: env.GetDefault("APPVEYOR_REPO_COMMIT", ""),
	Job:    env.GetDefault("APPVEYOR_JOB_BRANCH", ""),
	PR:     env.GetDefault("APPVEYOR_PULL_REQUEST_NUMBER", ""),
	Slug:   env.GetDefault("APPVEYOR_REPO_NAME", ""),
	Root:   env.GetDefault("APPVEYOR_BUILD_FOLDER", ""),
}

// Bamboo Service - https://confluence.atlassian.com/bamboo/bamboo-variables-289277087.html
var bamboo = &Environment{
	Base:   "bamboo_agentId",
	Build:  env.GetDefault("bamboo_buildNumber", ""),
	Branch: env.GetDefault("bamboo_planRepository_1_branchName", ""),
	Commit: env.GetDefault("bamboo_planRepository_1_revision", ""),
	Job:    env.GetDefault("bamboo_buildKey", ""),
	PR:     "",
	Root:   env.GetDefault("bamboo_build_working_directory", ""),
	Slug:   "",
}

// BuildKite Service - https://buildkite.com/docs/builds/environment-variables<Paste>
var buildKite = &Environment{
	Base:   "BUILDKITE",
	Build:  env.GetDefault("BUILDKITE_BUILD_NUMBER", ""),
	Branch: env.GetDefault("BUILDKITE_BRANCH", ""),
	Commit: env.GetDefault("BUILDKITE_COMMIT", ""),
	PR:     pullRequestIsFalse("BUILDKITE_PULL_REQUEST"),
	Slug:   env.GetDefault("BUILDKITE_ORGANIZATION_SLUG", "") + "/" + env.GetDefault("BUILDKITE_PROJECT_SLUG", ""),
	Root:   env.GetDefault("BUILDKITE_BUILD_CHECKOUT_PATH", ""),
}

// CircleCI Service - https://circleci.com/docs/1.0/environment-variables
var circleCI = &Environment{
	Base:   "CIRCLECI",
	Build:  env.GetDefault("CIRCLE_BUILD_NUM", "") + "." + env.GetDefault("CIRCLE_NODE_INDEX", ""),
	Branch: env.GetDefault("CIRCLE_BRANCH", ""),
	Commit: env.GetDefault("CIRCLE_SHA1", ""),
	Job:    env.GetDefault("CIRCLE_BUILD_NUM", "") + "." + env.GetDefault("CIRCLE_NODE_INDEX", ""),
	PR: func() string {
		if pr := env.GetDefault("CI_PULL_REQUEST", ""); pr != "" {
			tmp := strings.Split(pr, "/")
			return tmp[len(tmp)-1]
		}

		return ""
	}(),
	Slug: env.GetDefault("CIRCLE_PROJECT_USERNAME", "") + "/" + env.GetDefault("CIRCLE_PROJECT_REPONAME", ""),
	Root: "",
}

// Codeship Service - https://documentation.codeship.com/basic/builds-and-configuration/set-environment-variables
var codeship = &Environment{
	// Base == "codeship"
	Base:   "CI_NAME",
	Build:  env.GetDefault("CI_BUILD_NUMBER", ""),
	Branch: env.GetDefault("CI_BRANCH", ""),
	Commit: env.GetDefault("CI_COMMIT_ID", ""),
	Job:    "",
	PR:     "",
	Slug:   env.GetDefault("CI_REPO_NAME", ""),
	Root:   "",
}

// Drone Service - http://readme.drone.io/0.5/usage/environment-reference
var drone = &Environment{
	Base:   "DRONE",
	Build:  env.GetDefault("DRONE_BUILD_NUMBER", ""),
	Branch: env.GetDefault("DRONE_BRANCH", ""),
	Commit: env.GetDefault("DRONE_COMMIT_SHA", ""),
	Job:    env.GetDefault("DRONE_JOB_NUMBER", ""),
	PR:     env.GetDefault("DRONE_PULL_REQUEST", ""),
	Slug:   env.GetDefault("DRONE_REPO_OWNER", "") + "/" + env.GetDefault("DRONE_REPO_NAME", ""),
	Root:   "",
}

// Gitlab Service - https://docs.gitlab.com/ce/ci/variables/README.html
var gitlab = &Environment{
	Base:   "DRONE",
	Build:  env.GetDefault("CI_JOB_NAME", ""),
	Branch: env.GetDefault("CI_COMMIT_REF_NAME", ""),
	Commit: env.GetDefault("CI_COMMIT_SHA", ""),
	Job:    env.GetDefault("CI_JOB_STAGE", ""),
	PR:     "",
	Slug: func() string {
		if slug := env.GetDefault("CI_REPOSITORY_URL", ""); slug != "" {
			tmp := strings.Split(slug, "/")
			tmp = tmp[3:5]
			slug = strings.Join(tmp, "/")
			slug = strings.Replace(slug, ".git", "", 1)
			return slug
		}

		return ""
	}(),
	Root: env.GetDefault("CI_PROJECT_DIR", ""),
}

// Jenkins Service - https://wiki.jenkins.io/display/JENKINS/Building+a+software+project
var jenkins = &Environment{
	Base:   "JENKINS_URL",
	Build:  env.GetDefault("BUILD_NUMBER", ""),
	Branch: env.GetDefault("GIT_BRANCH", ""),
	Commit: env.GetDefault("GIT_COMMIT", ""),
	Job:    "",
	PR:     "",
	Slug:   "",
	Root:   env.GetDefault("WORKSPACE", ""),
}

// Semaphore Service - https://semaphoreci.com/docs/available-environment-variables.html
var semaphore = &Environment{
	Base:   "SEMAPHORE",
	Build:  env.GetDefault("SEMAPHORE_BUILD_NUMBER", ""),
	Branch: env.GetDefault("BRANCH_NAME", ""),
	Commit: "",
	Job:    "",
	PR:     env.GetDefault("PULL_REQUEST_NUMBER", ""),
	Slug:   env.GetDefault("SEMAPHORE_REPO_SLUG", ""),
	Root:   env.GetDefault("SEMAPHORE_PROJECT_DIR", ""),
}

// Shippable Service - http://docs.shippable.com/ci/env-vars/#stdEnv
var shippable = &Environment{
	Base:  "SHIPPABLE",
	Build: env.GetDefault("BUILD_NUMBER", ""),
	Branch: func() string {
		if branch := env.GetDefault("BASE_BRANCH", ""); branch != "" {
			return branch
		}

		return env.GetDefault("Branch", "")
	}(),
	Commit: "",
	Job:    env.GetDefault("JOB_NUMBER", ""),
	PR:     pullRequestIsFalse("PULL_REQUEST"),
	Slug:   env.GetDefault("SHIPPABLE_REPO_SLUG", ""),
	Root:   env.GetDefault("SHIPPABLE_BUILD_DIR", ""),
}

// Teamcity Service - https://confluence.jetbrains.com/display/TCD10/Predefined+Build+Parameters
var teamcity = &Environment{
	Base:   "TEAMCITY_VERSION",
	Build:  env.GetDefault("BUILD_NUMBER", ""),
	Branch: "", // TODO
	Commit: env.GetDefault("BUILD_VCS_NUMBER", ""),
	Job:    "",
	PR:     "",
	Slug:   env.GetDefault("TEAMCITY_BUILDCONF_NAME", ""),
	Root:   "", // TODO
}

// Travis Service - https://docs.travis-ci.com/user/environment-variables
var travis = &Environment{
	Base:   "TRAVIS",
	Build:  env.GetDefault("TRAVIS_BUILD_NUMBER", ""),
	Branch: env.GetDefault("TRAVIS_BRANCH", ""),
	Commit: env.GetDefault("TRAVIS_COMMIT", ""),
	Job:    env.GetDefault("TRAVIS_JOB_NUMBER", ""),
	PR:     pullRequestIsFalse("TRAVIS_PULL_REQUEST"),
	Slug:   env.GetDefault("TRAVIS_REPO_SLUG", ""),
	Root:   env.GetDefault("TRAVIS_BUILD_DIR", ""),
}

// Wercker Service - http://devcenter.wercker.com/docs/environment-variables/available-env-vars#hs_cos_wrapper_name
var wercker = &Environment{
	Base:   "WERCKER_MAIN_PIPELINE_STARTED",
	Build:  env.GetDefault("WERCKER_MAIN_PIPELINE_STARTED", ""),
	Branch: env.GetDefault("WERCKER_GIT_BRANCH", ""),
	Commit: env.GetDefault("WERCKER_GIT_COMMIT", ""),
	Job:    "",
	PR:     "",
	Slug:   env.GetDefault("WERCKER_GIT_OWNER", "") + "/" + env.GetDefault("WERCKER_GIT_REPOSITORY", ""),
	Root:   env.GetDefault("WERCKER_ROOT", ""),
}
