# Storyblok Plugin for Mach Composer 

This repository contains the Storyblok plugin for Mach Composer. It requires Mach Composer 3.x

It uses the Terraform provider for Storyblok, see https://github.com/labd/terraform-provider-storyblok/

## Usage

```yaml
mach_composer:
  plugins:
    storyblok:
      source: mach-composer/storyblok
      version: 0.1.0
      
global:
  # ...
  storyblok:
    url: management-api
    token: your-client-token
  
sites:
  - identifier: my-site
    # ...
    storyblok:
      space_id: your-hub-id
      url: management-api # override
      token: your-client-token # override
```
