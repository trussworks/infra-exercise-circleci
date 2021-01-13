# Truss Engineering Playbook InfraSec CircleCI exercise(s)

A really empty repo to demonstrate how to setup a repo's CI.
For great justice.

To complete this exercise you will need to have basic experience with:

* command line
* building and testing your code
* git and GitHub

Please refer to the following pieces of documentation for additional philosophy or reference:

* [Engineering Playbook Intro to CI/CD](https://github.com/trussworks/Engineering-Playbook/blob/main/developing/cicd/intro.md)
* [Engineering Playbook CircleCI common patterns](https://github.com/trussworks/Engineering-Playbook/blob/main/developing/cicd/circleci-patterns.md)
* [12 Factor App: Build Release Run](https://12factor.net/build-release-run)
* [Google SRE book: Ch 8 Release Engineering](https://sre.google/sre-book/release-engineering/)

## Exercise 1: Set up a basic CI flow in CircleCI

Fork this repo.
Alternatively, if you have a project you would like to add CI to make sure it builds and has at least one test.

Then answer the following questions:

* How would you build this code locally?
* How would you test this code locally?
* How would you like to distribute this code?
* How would you like to deploy this code?
* _Bonus:_ How are you going to be versioning this code?
* _Bonus:_ How are you going to configure this code when you deploy it?

### I. Define your CI pipelines in CircleCI config

Conceptually, you will have 3 major jobs in your basic integration pipeline:

* build
* test
* deploy

Define what steps you would need to take for each job. Use your answers from above to define the commands you want your jobs to run.

To get you started on understanding CircleCI configuration, there's some basic vocabulary we'll need to go over:

* __executor__ - The computer that is running your defined CI configuration. This can be a whole VM or a Docker container.
* __job__ - A unit of work that you need accomplished. These can be made up of multiple steps. Jobs can be reused in multiple workflows.
* __step__ - A step in the associated job. Generally one set of commands.
* __workflow__ - A series of jobs that can be run in sequence or in parallel to completion. This can be just one job.
* __context__ - Each individual workflow/job/step can have different context that can be configured in a number of ways. This may be files on disk, environment variables, or even your commit hash. We will _not_ go over the details here but leave this for a later exercise.

If you want more info or a general reference, see the full [CircleCI config docs](https://circleci.com/docs/2.0/config-intro/#section=configuration).

Using the [template](.circleci/config.yml) start to fill in steps for the jobs corresponding to `build`, `test`, and `deploy`.

TODO: Add a walkthrough of the template.yml file?

Save your new configuration file to `.circleci/config.yml` and commit it to your git repo and push to GitHub.

_Note:_ It does but it doesn't matter which branch you push to. Once configured, CircleCI will attempt to run the configuration on every push and PR.

#### __How do I know my config is right?__

Honestly, you could just push and continue with configuring the project but we highly recommend the CircleCI command line utility to make this faster.

To install on your Mac with Homebrew run:

```sh
brew install circleci
```

To validate your configuration run somewhere in your repo:

```sh
circleci config validate
```

The output can be hard to read but will at least give you a sense of whether CircleCI can run it.

### II. Integrate Github With CircleCi

Log into CircleCI in your web browser. If you installed the CircleCI CLI you can run:

```sh
circleci open
```

to go directly to the CircleCI dashboard for this project.

TODO: write down what I do here...

After being configured, CircleCI will automatically run your configured workflow(s).

Open up the GitHub page for your project.

TODO: write down what I do here...

## Exercise 2: More CircleCi config

Topics to cover:

TODOS:

* Configure Caching
* Injecting secrets safely
* CircleCI orbs
* Other executors
* CircleCI contexts
* How to reuse config more effectively

## Exercise 3: Improve flows

Topics to cover:

TODOS:

* Release based on version strings
* Manual approvals
