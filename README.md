# Como utilizar ?

- Envie um requisição GET para https://gofrete.herokuapp.com/frete
- Os parâmetros que devem ser enviados devem possuir os nomes obrigatóriamente iguais ao exemplo abaixo:

CdServico=40010 (Obs: 40010 -> SEDEX Varejo,40045 -> SEDEX a Cobrar Varejo,40215 -> SEDEX 10 Varejo,40290 -> SEDEX Hoje Varejo,41106 -> PAC Varejo)

CepOrigem=11111111 (Obs: Utilize um CEP sem caracteres especiais)

CepDestino=11111111 (Obs: Utilize um CEP sem caracteres especiais)

VlPeso=0.100 (Obs: O correios utiliza o formato de gramas para calcular o peso.)

CdFormato=1

VlComprimento=11

VlAltura=11

VlLargura=11

VlDiametro=11

CdMaoPropria=S

VlValorDeclarado=100

CdAvisoRecebimento=S
