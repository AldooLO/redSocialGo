package main

import (
	"context"
	"os"
	"redSocialGo/awsgo"
	"redSocialGo/bd"
	"redSocialGo/handlers"
	"redSocialGo/models"
	"redSocialGo/secretmanager"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse
	awsgo.InicializoAWS()

	if !ValidoParametros() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno. deben incluir 'SecretName', 'BucketName', 'UrlPrefix'",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura de Secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	path := strings.Replace(request.PathParameters["redSocial"], os.Getenv("UrlPrefix"), "", -1)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), request.HTTPMethod)

	err = bd.ConectarBD(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error conectando la BD " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	respAPI := handlers.Manejadores(awsgo.Ctx, request)
	if respAPI.CustomResp == nil{
		res = &events.APIGatewayProxyResponse{
			StatusCode: respAPI.Status,
			Body:       respAPI.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}else{
		return respAPI.CustomResp, nil
	}
}

func ValidoParametros() bool {
	_, traeParametro := os.LookupEnv("SecretName")

	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("BucketName")

	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("UrlPrefix")

	if !traeParametro {
		return traeParametro
	}

	return traeParametro
}
