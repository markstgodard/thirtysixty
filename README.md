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
  - https://www.newegg.ca/zotac-geforce-rtx-3060-ti-zt-a30610h-10m/p/N82E16814500507?Description=3060%20ti&cm_re=3060_ti-_-14-500-507-_-Product
  - https://www.newegg.ca/zotac-geforce-rtx-3060-ti-zt-a30610e-10m/p/N82E16814500506?Description=3060%20ti&cm_re=3060_ti-_-14-500-506-_-Product
  - https://www.newegg.ca/msi-geforce-rtx-3060-ti-rtx-3060-ti-ventus-2x-oc/p/N82E16814137612?Description=3060%20ti&cm_re=3060_ti-_-14-137-612-_-Product
  - https://www.newegg.ca/asus-geforce-rtx-3060-ti-tuf-rtx3060ti-o8g-gaming/p/N82E16814126471?Description=3060%20ti&cm_re=3060_ti-_-14-126-471-_-Product
  - https://www.newegg.ca/msi-geforce-rtx-3060-ti-rtx-3060-ti-gaming-x-trio/p/N82E16814137611?Description=3060%20ti&cm_re=3060_ti-_-14-137-611-_-Product
  - https://www.newegg.ca/evga-geforce-rtx-3060-ti-08g-p5-3663-kr/p/N82E16814487535?Description=3060%20ti&cm_re=3060_ti-_-14-487-535-_-Product
  - https://www.newegg.ca/gigabyte-geforce-rtx-3060-ti-gv-n306teagle-8gd/p/N82E16814932379?Description=3060%20ti&cm_re=3060_ti-_-14-932-379-_-Product
```
