# Logentries Provider

The provider configuration block accepts the following arguments:

* ``account_key`` - Your Account Key for the the Logentries API. Find it per instructions at https://logentries.com/doc/accountkey/. May alternatively be set via the ``LOGENTRIES_KEY`` environment variable.


## Example Usage

```terraform
provider "logentries" {
    account_key = "12345678-1234-1234-5678-901234567890"
}

resource "logentries_logset" "my_host" {
  name = "MyHost"
}

resource "logentries_log" "my_log" {
  name = "terraform_test_log"
  in_logset_with_key = "${logentries_logset.my_host.key}"
}
```

## Contributing

How to submit changes:

1. Fork this repository.
2. Make your changes.
3. Email us at opensource@tulip.co to sign a CLA.
4. Submit a pull request.


## Who's Behind It

terraform-provider-logentries is maintained by Tulip. We're an MIT startup located in Boston, helping enterprises manage, understand, and improve their manufacturing operations. We bring our customers modern web-native user experiences to the challenging world of manufacturing, currently dominated by ancient enterprise IT technology. We work on Meteor web apps, embedded software, computer vision, and anything else we can use to introduce digital transformation to the world of manufacturing. If these sound like interesting problems to you, [we should talk](mailto:jobs@tulip.co).


## License

terraform-provider-logentries is licensed under the [Apache Public License](LICENSE).
