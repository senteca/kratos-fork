import {
  extractRecoveryCode,
  gen,
  MOBILE_URL,
  website
} from '../../../../helpers'

context('Mobile Profile', () => {
  describe('Recovery Flow Errors', () => {
    before(() => {
      cy.enableRecovery('code')
      cy.useConfigProfile('mobile')
    })

    let email, password

    before(() => {
      email = gen.email()
      password = gen.password()
      cy.registerApi({ email, password, fields: { 'traits.website': website } })
    })

    describe('code', () => {
      beforeEach(() => {
        cy.deleteMail()
        cy.longCodeLifespan()
      })

      it('fails with validation errors', () => {
        cy.visit(MOBILE_URL + '/Recovery')
        cy.get('*[data-testid="field/email"] input[data-testid="email"]')
          .clear()
          .type('not-an-email')
        cy.get(
          '*[data-testid="field/method/code"] div[data-testid="submit-form"]'
        ).click()

        cy.get('*[data-testid="field/email"]').should(
          'contain.text',
          'is not valid "email"'
        )
      })

      it('shows code expired message if expired code is submitted', () => {
        cy.shortCodeLifespan()
        cy.visit(MOBILE_URL + '/Recovery')

        cy.get('*[data-testid="field/email"] input[data-testid="email"]').type(
          email
        )
        cy.get(
          '*[data-testid="field/method/code"] div[data-testid="submit-form"]'
        ).click()

        cy.get('*[data-testid="ui/message/1060003"]').should(
          'contain.text',
          'An email containing a recovery code has been sent to the email address you provided.'
        )

        cy.getMail().should((message) => {
          expect(message.subject).to.equal('Recover access to your account')
          expect(message.toAddresses[0].trim()).to.equal(email)

          const code = extractRecoveryCode(message.body)
          expect(code).to.not.be.undefined
          expect(code.length).to.equal(8)

          cy.get('*[data-testid="field/code"] input[data-testid="code"]').type(
            code
          )
        })
        cy.get(
          '*[data-testid="field/method/code"] div[data-testid="submit-form"]'
        ).click()

        cy.get('[data-testid="ui/message/4060005"]').should(
          'contain.text',
          'The recovery flow expired'
        )
      })

      it('fails on invalid code', () => {
        cy.visit(MOBILE_URL + '/Recovery')
        cy.get('*[data-testid="field/email"] input[data-testid="email"]')
          .clear()
          .type(email)
        cy.get(
          '*[data-testid="field/method/code"] div[data-testid="submit-form"]'
        ).click()

        cy.get('*[data-testid="ui/message/1060003"]').should(
          'contain.text',
          'An email containing a recovery code has been sent to the email address you provided.'
        )

        cy.get('*[data-testid="field/code"] input[data-testid="code"]').type(
          '123456'
        )

        cy.get(
          '*[data-testid="field/method/code"] div[data-testid="submit-form"]'
        ).click()

        cy.get('[data-testid="ui/message/4060006"]').should(
          'contain.text',
          'The recovery code is invalid or has already been used. Please try again.'
        )
      })
    })
  })
})
