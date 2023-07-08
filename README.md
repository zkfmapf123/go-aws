# go aws tools

## Tools

- json-parser
- generic

## aws

- set IAM -> set EC2

## tls

> self signed ca

- 자체 서명된 CA를 구성하여, 일반적으로 공용 웹사이트나 서비스가 아닌 내부 Test용으로 사용
- 자체 서명되기 때문에, 유효성부분이나 신뢰성 문제가 있다.

> Root Signed CA

- 인증된 CA 에서 발급한 인증서를 사용하여 통신 간의 신뢰를 구축한다.

## Reference

- <a href="https://aws.github.io/aws-sdk-go-v2/docs/getting-started/">go aws sdk document </a>
- <a href="https://aws.github.io/aws-sdk-go-v2/docs/code-examples/">go-aws-sdk example</a>
