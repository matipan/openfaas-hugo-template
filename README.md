# OpenFaaS template for Hugo
This template allows you to deploy hugo sites using OpenFaaS. It simply copies over the contents of your hugo site, builds it into the `public` directory and then uses a [very lightweight static server](https://gitlab.com/matipan/static-server) that serves the content and provides a healtcheck that follows the standards from OpenFaaS.

## Usage
Create the hugo function:
```sh
git init
faas template pull https://github.com/matipan/openfaas-hugo-template
faas new --lang hugo -g <openfaas gateway url> --prefix <docker hub username> example-site
```
This will create a folder called `example-site`, `cd` into it and now create the site with this instructions from the [hugo quickstart guide](https://gohugo.io/getting-started/quick-start/#step-2-create-a-new-site):
```sh
hugo new site .
git submodule add https://github.com/budparr/gohugo-theme-ananke.git themes/ananke
echo 'theme = "ananke"' >> config.toml
```
At this point you can run `hugo server` in that directory to build and test your site locally without needing to deploy it to OpenFaaS. Remember to update the `baseURL` found in the `config.toml` to the domain that you will be using.

## Custom domains with TLS on OpenFaaS
Follow [this tutorial](https://blog.matiaspan.dev/posts/bringing-serverless-to-a-webpage-near-you-with-hugo-and-kubernetes/) to learn how to continuosly deploy your sites and have a custom domain name pointing at them (with TLS!).
