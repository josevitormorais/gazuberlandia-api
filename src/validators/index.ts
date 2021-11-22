import * as Yup from "yup"

const signInValidateSchema = Yup.object().shape({
  email: Yup.string().email().required(),
  password: Yup.string().min(4).required()
})

export { signInValidateSchema }
