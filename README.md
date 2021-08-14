# LiveDoc
Um sistema para gerenciamento de propostas comerciais

## Requisitos Funcionais

### Cadastrar um novo cliente de prospecção

- CNPJ
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
  - Valor faturamento anual
  - Quantidade estimada de emissão de documentos por tipo de documento 
    - Quantidade NF-e
    - Quantidade NFS-e
    - Quantidade CT-e
    - Quantidade NFC-e
  - NF-e
    - Em quais estados a empresa emite NF-e
  - NFS-e
    - Em quais municípios a empresa emite NFS-e?
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
      - Nota fiscal standard SAP está implementada?
        - Se sim
          - BADI CL_NFE_PRINT já implementada?
          - Layout de DANFE já desenvolvido?
          - Possui alguma customização no processo de envio?
      - Qual o tipo de comunicação deseja utilizar?
        - Comunicação HTTP direto do SAP ECC ou S/4HANA
        - CPI
        - PI
        - PO
        - Outra?
          - Informar
        - Nível de autorizações / acessos na solução de comunicação?
          - Nenhum
          - Controle extremo
          - Controle moderado
          - Acesso Completo
    - Outro
      - Informar

  - Existirá alguma integração do Orbit com outra aplicação?
  
\* Mapeado como ponto de melhoria futura!