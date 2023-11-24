# **Code Runner software system - Project plan.**

Author: Yordan Mitev

Stakeholder: Yordan Mitev

## **Problem Statement**

Currently, there is a lack of a unified solution for automated code execution and
grading. There are environments for running code, but there are no plug-and-play
versions of a tool of this variety. As a student it is awkward to see an
institution's classes that teach (basic) algorithms and language basics to not
have an established process for verifying solutions and if your code doesn't
function properly, reach out to a teacher, and have a review consultation/
meeting. For example, when teaching SQL, there is a system to run and verify
queries, but it is very janky and not very well documented

### Why you shouldn't use a SaaS platform?

When we talk about a SaaS platform like HackerRank and Codechef, we accept the
premise that we need to leave our sovereign space where we control who is
registered and has access to our resources.
[Why selfhosting is important in this day and age.](https://dataswamp.org/~solene/2021-07-23-why-selfhosting-is-important.html)

- <span style="color:red">**Drawbacks**</span>:
  - **Proprietary Solutions**: SaaS platforms often use proprietary solutions
    that are not publicly accessible. This limits the ability to customize
    orintegrate these solutions into external systems.
  - **Full-scale services**: As it stands, these SaaS platforms also base their
    business model on the fact that they provide a predefined set of exercises,
    ranked by difficulty and are fundamentally incompatible with a self-hosted
    platform as the goal is.

## **Research**

### [Code running in depth](code-running.md)

### TLDR

Currently, there exist commercial options like hackerrank, and codewars, but they
are not on-premise, are closed-source, and allow not so much customisation. They
are good if you have a deep budget and primarily focus on interviews and want
access to a lot of predefined challenges.

There is a platform called Open Judge System which has been maintained by a
Bulgarian company, but has not received any features/updates/new languages for
over 8 years. It is written in .NET Framework C#, and looks interesting, but its
execution engine is really hard to maintain, OS-dependent and hard to configure/
scale.

There is an open-source code execution engine called "Piston" that allows
on-premises hosting and code execution over an HTTP api. It has support for a
wide variety of languages, and it is very simple to change to a specific language
version, or support a new language. There seems to exist support for SQL, but is
rather minimal.

There is another tool called Judge0 that offers both on-premises hosting and
paid-for hosting, but it includes a lot of features that are not needed for a
project of this caliber including:

- Batched code execution
- Authentication
- Run history

To add to that, OpenJudgeSystem and Judge0 are both GPL-3 licensed software,
which will conflict with a MIT-Licensed version.

## **Technology**

The project will be developed in Rust (programming language), because I want to
further my understanding of the language. Though there are more fitting options,
I want to explore more about the language and its features.

When it comes to making UIs, there are several interesting approaches that should
be considered:

- Server-side-rendered templates:
  - Jinja
  - Handlebars

- Javascript frameworks:
  - Svelte
  - SolidJS

## **Deliverables**

Application hosted on a server + API

Demo-able tool for running receiving code

Integration path of the API (canvaslms extension or fontys oauth)

Modified Piston runtime.

## **Timeline**

### **Sprint 1: (until: 13th Oct 2023)**

- Technology picks
- Hosting location (VPS, self-hosted?)
- Repository setup

### **Sprint 2: (until: 3rd Nov 2023)**

- Basic data types definition and functionality (Users, Problem, Problem
  collections, Solutions, etc.)
- Basic authentication (username-password based)
- Basic database configuration

### **Sprint 3: (until: 24th Nov 2023)**

- Connect the execution engine to the application
- Consider advanced authentication options (canvaslms/fontys oauth)

### **Sprint 4: (until: 22nd Dec 2023)**

- Extend the execution engine to include more build system configurations
- SQL query runner
- Implement one of the integration paths from **Sprint 3**

### **Sprint 5: TBD**
