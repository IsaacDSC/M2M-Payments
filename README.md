# M2M-Payments


######

```sh
mockgen -source=./<path> -destination=./<path>/<filename>.mock.go -package=mock
```

### 
Qual Arquitetura utilizar?

Expecifícações:
  - Projeto tende a ter funcionalizades novas e mudanças frequentes
  - Projeto não pode estar amarrado a infraestrutura podendo haver mudanças
  - Independente de framework web
  - Necessário logs e seria bom ter independencia de loggers podendo ser alterada
  - Regra de negócio pode mudar de acordo com a empresa e as regras de pci
  - Projeto precisa ter observabilidade
  - Projeto precisa ter no minimo de 80% de cobertura
  - Haverá um time para cuidar do projeto inicialmente 3 pessoas