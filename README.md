# LiveDoc
Um sistema para gerenciamento de propostas comerciais

## Requisitos Funcionais

### Cadastrar um novo cliente de prospecção

- CNPJ
- Razão Social
- Inscrição Estadual
- Inscrição Municipal
- Endereço
- Número
- Bairro
- Cidade
- UF
- Complemento

### Listar os clientes cadastrados

### Criar um novo questionário
- Selecionar o cliente de prospecção
- Selecionar o produto*
- Preencher o motivo
- Adicionar endereços de e-mail autorizados para preenchimento do formulário
  - Gerar e enviar senha para o e-mail cadastrado
- 

### Listar todos os questionários criados

### Acessar o questionário através de uma URL única

- Preencher e-mail e senha para acesso ao questionário
- Permitir preenchimento parcial e assíncrono
- Dados do questionário:
  - Idioma oficial do projeto
  - Quais sistemas emitirão nota?
    - SAP
      - Versão do SAP
      - Versão do componente SAP_BASIS
      - Versão do componente SAP_APPL
      - Nível de autonomia para criação de request?
        - Nenhuma autonomia
        - Controle extremo
        - Controle moderado
        - Autonomia total
      - Nível de autonomia para transporte de request?
        - Nenhuma autonomia
        - Controle extremo
        - Controle moderado
        - Autonomia total
    - Outro
      - Informar
  - Valor faturamento anual
  - Quantidade de NFs de saída por mês
  - Quantidade de NFs de entrada por mês
  - NF-e
    - Em quais estados a empresa emite NF-e
  - NFS-e
    - Em quais municípios a empresa emite NFS-e?
  - Qual o tipo de comunicação deseja utilizar?
    - Comunicação HTTP direto do SAP ECC ou S/4HANA
    - CPI
    - PI
    - PO
    - Outra?
      - Informar
    - Nível de autorizações / acessos?
        - Nenhum
        - Controle extremo
        - Controle moderado
        - Completo
  - Existirá alguma integração do Orbit com outra aplicação?


\* Mapeado como ponto de melhoria futura!