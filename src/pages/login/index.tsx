import React from "react"
import api from "../../services/api"
import NextLink from 'next/link'
import { Link } from '@chakra-ui/react'
import { Button } from "@chakra-ui/button"
import { Input } from "@chakra-ui/input"
import {
  FormControl,
  FormLabel,
  FormErrorMessage
} from "@chakra-ui/form-control"
import { Flex, Heading } from "@chakra-ui/layout"
import { yupResolver } from "@hookform/resolvers/yup/dist/yup"
import { SubmitHandler, useForm } from "react-hook-form"
import { signInValidateSchema } from "../../validators"
import { useToast } from "@chakra-ui/react"
import { ToastMsgGenericErr } from '../../constants/errors'
import { AxiosResponse } from "axios"

type SignInData = {
  email: string
  password: string
}

const Login = () => {

  const { handleSubmit, register, formState: { errors, isSubmitting } } = useForm<SignInData>({
    resolver: yupResolver(signInValidateSchema),
    mode: "onBlur"
  })

  const toast = useToast()
  const toastId = 'has-signIn-tost'
  
  const handleSignIn: SubmitHandler<SignInData> = (data) => api
      .post("/login", data)
      .then(({ data }: AxiosResponse) => console.log(data)) //TODO initially should store the token in localStorage, this can be a hook for example "useLocalStorage"
    .catch((err) => {
        if (!toast.isActive(toastId)) {
          toast({
            id: toastId,
            description: (err.response?.data?.message || ToastMsgGenericErr),
            status: "info",
            duration: 3000,
            isClosable: true,
            position: "top-right"
          })
        }
    })

  return (
    <Flex height="100vh" alignItems="center" justifyContent="center">
      <Flex flexDirection="column" alignItems="center" justifyContent="center" p={12} rounded={6}>
        <Heading mb={6}> Faça seu Login</Heading>
        <form onSubmit={handleSubmit(handleSignIn)}>
          <Flex flexDirection="column" alignItems="center" justifyContent="center" p={6} rounded={6}>
            <FormControl isInvalid={Boolean(errors.email?.message)}>
              <FormLabel htmlFor="email">E-mail</FormLabel>
              <Input
                id="email"
                type="email"
                placeholder="yourmail@example.com"
                {...register("email")}
              />
              <FormErrorMessage mb={4}>
                {errors.email && errors.email?.message}
              </FormErrorMessage>
            </FormControl>

            <FormControl isInvalid={Boolean(errors.password?.message)}>
              <FormLabel htmlFor="password">Password</FormLabel>
              <Input
                id="password"
                type="password"
                placeholder="password"
                {...register("password")}
              />
              <FormErrorMessage mb={4}>
                {errors.password && errors.password?.message}
              </FormErrorMessage>
            </FormControl>
            <Button mt={5} isLoading={isSubmitting} type="submit">
              Continuar
            </Button>
          </Flex>
        </form>

        <NextLink href="/register">
          <Link mb={3}>Não sou cliente</Link>
        </NextLink>
        <NextLink href="/forgot" >
          <Link>Esqueci minha senha</Link>
        </NextLink>

      </Flex>
    </Flex>
  )
}

export default Login
