<!--- This file is autogenerated from the files in docsgenerator/templates/configure-director --->
&larr; [back to Commands](../README.md)

# `om configure-director`

The `configure-director` command will allow you to setup your BOSH tile on the OpsManager.

Once configured, changes will not take affect until the next [`apply-changes`](../apply-changes/README.md).

## Supported Infrastructures
* [AWS](aws.md)
* [GCP](gcp.md)
* [Azure](azure.md)
* [vSphere](vsphere.md)

## Command Usage
```
Usage:
  om [OPTIONS] configure-director [configure-director-OPTIONS]

This authenticated command configures the director.

Application Options:
      --ca-cert=                      OpsManager CA certificate path or value
                                      [$OM_CA_CERT]
  -c, --client-id=                    Client ID for the Ops Manager VM (not
                                      required for unauthenticated commands)
                                      [$OM_CLIENT_ID]
  -s, --client-secret=                Client Secret for the Ops Manager VM (not
                                      required for unauthenticated commands)
                                      [$OM_CLIENT_SECRET]
  -o, --connect-timeout=              timeout in seconds to make TCP
                                      connections (default: 10)
                                      [$OM_CONNECT_TIMEOUT]
  -d, --decryption-passphrase=        Passphrase to decrypt the installation if
                                      the Ops Manager VM has been rebooted
                                      (optional for most commands)
                                      [$OM_DECRYPTION_PASSPHRASE]
  -e, --env=                          env file with login credentials
  -p, --password=                     admin password for the Ops Manager VM
                                      (not required for unauthenticated
                                      commands) [$OM_PASSWORD]
  -r, --request-timeout=              timeout in seconds for HTTP requests to
                                      Ops Manager (default: 1800)
                                      [$OM_REQUEST_TIMEOUT]
  -k, --skip-ssl-validation           skip ssl certificate validation during
                                      http requests [$OM_SKIP_SSL_VALIDATION]
  -t, --target=                       location of the Ops Manager VM
                                      [$OM_TARGET]
      --trace                         prints HTTP requests and response
                                      payloads [$OM_TRACE]
  -u, --username=                     admin username for the Ops Manager VM
                                      (not required for unauthenticated
                                      commands) [$OM_USERNAME]
      --vars-env=                     load vars from environment variables by
                                      specifying a prefix (e.g.: 'MY' to load
                                      MY_var=value) [$OM_VARS_ENV]
  -v, --version                       prints the om release version

Help Options:
  -h, --help                          Show this help message

[configure-director command options]
          --ignore-verifier-warnings  option to ignore verifier warnings. NOT
                                      RECOMMENDED UNLESS DISABLED IN OPS MANAGER
      -c, --config=                   path to yml file containing all config
                                      fields (see
                                      docs/configure-director/README.md for
                                      format)
      -l, --vars-file=                load variables from a YAML file
          --vars-env=                 load variables from environment variables
                                      (e.g.: 'MY' to load MY_var=value)
                                      [$OM_VARS_ENV]
      -v, --var=                      load variable from the command line.
                                      Format: VAR=VAL
          --ops-file=                 YAML operations file
```

### Configuring via file

The `--config` flag is available for convenience to allow you to pass a single
file with all the configuration required to configure your director.

When providing a single config file each of the other individual flags maps to a
top-level element in the YAML file.

#### Example YAML:
```yaml
---
az-configuration:
- name: some-az
director-configuration:
  max_threads: 5
iaas-configuration:
  iaas_specific_key: some-value
network-assignment:
  network:
    name: some-network
networks-configuration:
  networks:
  - network: network-1
resource-configuration:
  compilation:
    instance_type:
      id: m4.xlarge
security-configuration:
  trusted_certificates: some-certificate
syslog-configuration:
  syslogconfig: awesome
vmextensions-configuration:
- name: a_vm_extension
  cloud_properties:
    source_dest_check: false
- name: another_vm_extension
  cloud_properties:
    foo: bar
vmtypes-configuration:
  custom_only: false
  vm_types:
  - name: a1.large
    cpu: 4
    ram: 8192
    ephemeral_disk: 10240
  - name: t2.small
    cpu: 1
    ram: 512
    ephemeral_disk: 1024
```

#### vmtypes-configuration:

Will set or update custom VM types on the director. If `custom_only` is `true`, 
the VM types specified in your configuration will be the **entire** list of
available VM types in the Ops Manager. If `false` or omitted, it will add the 
listed VM types to the list of default VM types for your IaaS. If a specified
VM type is named the same as a predefined VM type, it will overwrite the predefined
type. If multiple specified VM types have the same name, the one specified last
will be created. In either case, existing custom VM types do not persist across
`configure-director` calls, and it should be expected that the entire list of custom
VM types is specified in the director configuration.

##### Minimal example
```yaml
vmtypes-configuration:
  custom_only: false
  vm_types:
  - name: x1.large
    cpu: 8
    ram: 8192
    ephemeral_disk: 10240
  - name: mycustomvmtype
    cpu: 4
    ram: 16384
    ephemeral_disk: 4096
```

#### Variables

The `configure-director` command now supports variable substitution inside the config template:

```yaml
# config.yml
network-assignment:
  network:
    name: ((network_name))
```

Values can be provided from a separate variables yaml file (`--vars-file`) or from environment variables (`--vars-env`).

To load variables from a file use the `--vars-file` flag.

```yaml
# vars.yml
network_name: some-network
```

```
om configure-director \
  --config config.yml \
  --vars-file vars.yml
```

To load variables from a set of environment variables, specify the common
environment variable prefix with the `--vars-env` flag.

```
OM_VAR_network_name=some-network om configure-director \
  --config config.yml \
  --vars-env OM_VAR
```

The interpolation support is inspired by similar features in BOSH. You can
[refer to the BOSH documentation](https://bosh.io/docs/cli-int/) for details on how interpolation
is performed.