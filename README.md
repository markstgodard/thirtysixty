# thirtysixty

## About
Simple utility that fetches products and checks price and availability.

For any products that are found to be availble, will alert via SMS.


## Setup
Copy .envrc.example to .envrc and configure based on your Twilio account and numbers.

```bash
export TWILIO_ACCOUNT_SID="YOUR TWILIO ACCOUNT SID"
export TWILIO_AUTH_TOKEN="YOUR TWILIO AUTH TOKEN"
export TWILIO_FROM_PHONE_NUMBER="YOUR TWILIO PHONE NUMBER"
export TWILIO_TO_PHONE_NUMBER="YOUR PHONE NUMBER"
```


config.yml contains a list of products to get alerts on

```yaml
urls:
  - https://www.newegg.ca/asus-geforce-rtx-3060-ti-dual-rtx3060ti-o8g/p/N82E16814126468?Description=3060%20ti&cm_re=3060_ti-_-14-126-468-_-Product
  - https://www.newegg.ca/asus-geforce-rtx-3060-ti-ko-rtx3060ti-o8g-gaming/p/N82E16814126474?Description=3060%20ti&cm_re=3060_ti-_-14-126-474-_-Product
  - https://www.newegg.ca/evga-geforce-rtx-3060-ti-08g-p5-3667-kr/p/N82E16814487537?Description=3060%20ti&cm_re=3060_ti-_-14-487-537-_-Product
  - https://www.newegg.ca/asus-geforce-rtx-3060-ti-rog-strix-rtx3060ti-o8g-gaming/p/N82E16814126470?Description=3060%20ti&cm_re=3060_ti-_-14-126-470-_-Product
```
