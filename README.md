# Golang support for structured references

Package structref contains data types and validation functions for structured
references used on Swiss payment slips: creditor references as described in
ISO 11649, and Swiss ESR numbers. This is not an officially supported Google
product.

The types for different structured references all implement the Printer
interface defined in this package:

```
type Printer interface {
  DigitalFormat() string
  PrintFormat() string
}
```
