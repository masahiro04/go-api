package testData

var CompanyDefinition = `
"Company": {
  "type": "object",
  "properties": {
	"id": {
      "type": "string"
    },
    "name": {
      "type": "string"
    },
	"representativeLastName": {
      "type": "string"
    },
	"representativeFirstName": {
      "type": "string"
    },
	"representativeTitle": {
      "type": "string"
    },
	"zip": {
      "type": "string"
    },
	"city": {
      "type": "string"
    },
	"street": {
      "type": "string"
    },
	"building": {
      "type": "string"
    },
	"hubspotCompanyId": {
      "type": "string"
    },
 	"createdAt": {
      "type": "string",
      "format": "date-time"
    },
    "updatedAt": {
      "type": "string",
      "format": "date-time"
    }
  },
  "required": [
    "name",
	"createdAt",
    "updatedAt"
  ]
}`

var CompanyDeletedRespDefinition = `{
  "type": "object",
  "properties": {
    "company": {
      "type": "boolean"
	}
  }
}`

var CompanySingleRespDefinition = `{
  ` + CompanyDefinition + `,
  "type": "object",
  "properties": {
    "company": {
      "$ref": "#/Company"
    }
  },
  "required": [
    "company"
  ]
}`

var CompanyMultipleRespDefinition = `{
` + CompanyDefinition + `,
  "type": "object",
  "properties": {
    "companies": {
      "type": "array",
      "items": {
        "$ref": "#/Company"
      }
    }
  },
  "required": [
    "companies"
  ]
}`

var UserDefinition = `
"User": {
  "type": "object",
  "properties": {
	"Id": {
      "type": "integer"
    },
	"uuId": {
      "type": "string"
    },
	"companyId": {
      "type": "integer"
    },
    "name": {
      "type": "string"
    },
    "email": {
      "type": "string",
      "format": "email"
    },
	"kind": {
      "type": "string"
    },
	"role": {
      "type": "string"
    },
	"status": {
      "type": "string"
    },
	"hubspotContractId": {
      "type": "string"
    },
	"photoUrl": {
      "type": "string"
    },
	"signInCount": {
      "type": "integer"
    },
	"currentSignInAt": {
      "type": "string"
    },
	"lastSignInAt": {
      "type": "string"
    },
	"lastSignInIp": {
      "type": "string"
    },
	"currentSignInIp": {
      "type": "string"
    },
	"createdAt": {
      "type": "string",
		"format": "date-time"
    },
	"updatedAt": {
		"type": "string",
		"format": "date-time"
    },
	"deletedAt": {
      "type": "string",
		"format": "date-time"
    }
  },
  "required": [
	"Id",	
	"uuId",
	"companyId",
	"name",
	"email",
	"kind",
	"role",
	"status",
	"hubspotContractId",
	"signInCount",
	"currentSignInAt",
	"lastSignInAt",
	"lastSignInIp",
	"currentSignInIp",
	"createdAt",
	"updatedAt",
	"deletedAt"
  ]
}`

var UserRespDefinition = `{
	` + UserDefinition + `,
      "type": "object",
      "properties": {
        "user": {
          "$ref": "#/User"
        }
      },
	"required": [
    	"user"
  	]
}`

var ConfigResponse = `
"menuItem": {
  "type": "object",
  "properties": {
    "text": {
      "type": "string"
    },
	"path": {
      "type": "string"
    }
  },
  "required": [
    "text",
    "path"
  ]
}`

var MenuItemDefinition = `
"config": {
  "type": "object",
  "properties": {
	"headerMenu": {
      "type": "array",
      "items": {
        "$ref": "#/menuItem"
      }
    },
	"pullDownMenu": {
      "type": "array",
      "items": {
        "$ref": "#/menuItem"
      }
    }
  },
  "required": [
	"headerMenu",
	"pullDownMenu"
  ]
}`

var MeRespDefinition = `{
  ` + UserDefinition + `,` + MenuItemDefinition + `,` + ConfigResponse + `,
  "type": "object",
  "properties": {
    "user": {
      "$ref": "#/user"
    },
    "config": {
      "$ref": "#/config"
    }
  },
  "required": [
    "user",
    "config"
  ]
}`

var InvitationDefinition = `
"Invitation": {
  "type": "object",
  "properties": {
    "id": {
        "type": "string"
      },
    "companyID": {
      "type": "integer"
    },
	"invitedCompanyID": {
      "type": "int"
    },
	"isConfirmed": {
      "type": "boolean"
    },
 	"createdAt": {
      "type": "string",
      "format": "date-time"
    },
    "updatedAt": {
      "type": "string",
      "format": "date-time"
    }
  },
  "required": [
    "companyID",
    "isConfirmed",
    "invitedCompanyID",
	  "createdAt",
    "updatedAt"
  ]
}`

var InvitationRespDefinition = `{
  ` + InvitationDefinition + `,
  "type": "object",
  "properties": {
    "invitation": {
      "$ref": "#/Invitation"
    }
  },
  "required": [
    "invitation"
  ]
}`

var ErrorResponse = `{
  "type": "object",
  "properties": {
    "errors": {
      "type": "object",
      "properties": {
        "body": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      },
      "required": [
        "body"
      ]
    }
  },
  "required": [
    "errors"
  ]
}`
