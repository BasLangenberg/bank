# Bank

This is a software implementation of a bank. No real bank should use this as it is mosly made up by my own reasoning and an attempt to explore proper software engineering on saturday nights when I'm bored.

# Goals

This list is more a 'Bas is bored again and wants something to tickle his brain' then a promise this will eventually become feature complete like descibred below. Again, this is one of my many Saturday-Evenings-I-Am-Bored projects.

## Functional

Should be the focus instead of all technical mumbo jumbo.

- Open a bank account
- Wire money to another bank account
  - Figure out non-repudiation
- Get the amount of money in a bank account
- Get the transactions against a bank account
- Do transaction monitoring
- Do some sort of KYC from the banks perspective (Idea: Look into Camunda)
- Write OpenAPI specification for external clients to interact with
- Use JavaScript to build a real front-end

## Technical

- No merge to main unless all test pass
- Use a Test Driven Development (kinda) method to develop this piece of software
- Local tests should use Docker compose to test against real dependencies (Database, message queues...)
- Local CI/CD with Dagger + precommit hook
- Allow a merge to main only after succesfull CI (Github Runners)
- Automate releases on Github
- Include Grafana Cloud
  - Monitoring with Prometheus
  - Logs with Loki
  - Traces with Tempo
- Deployment target: Digital Ocean Kubernetes Engine
  - Terraform
  - ArgoCD
    - Use syncwaves

### Infra

- Keycloak (for auth)
- Postgres (for application data)

# Inspirations

Sources from which this repo (or its author) got inspirational insights, listed out in no particular order.

- [Ardanlabs/Service](https://github.com/Ardanlabs/Service)
- [LearnGo with tests](https://quii.gitbook.io/learn-go-with-tests)
- [Arnaud Lauret](http://apihandyman.io/) and his [book](https://www.manning.com/books/the-design-of-web-apis)
