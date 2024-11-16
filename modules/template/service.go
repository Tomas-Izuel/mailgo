package template

import "context"

func FindTemplateByIDService(templateId string,
	ctx context.Context) (*MailTemplate, error) {
	mailTemplate, err := findTemplateByID(templateId, ctx)
	if err != nil {
		return nil, err
	}

	return mailTemplate, nil
}

func createTemplateService(mailTemplate *CreateMailTemplateDto, ctx context.Context) (*MailTemplate, error) {
	id, err := createTemplate(mailTemplate, ctx)
	if err != nil {
		return nil, err
	}

	createdMailTemplate := &MailTemplate{
		ID:       id,
		Subject:  mailTemplate.Subject,
		BodyHTML: mailTemplate.BodyHTML,
	}

	return createdMailTemplate, nil
}
