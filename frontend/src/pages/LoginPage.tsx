import { useFormik } from "formik"
import * as Yup from "yup"
import { GoogleLogin } from "@react-oauth/google"

const LoginPage = () => {
    const formik = useFormik({
        initialValues:{
            usernameOrEmail : '',
            password: ''
        }, 
        validationSchema: Yup.object({
            usernameOrEmail: Yup.string().required("Username or email is required")
            .test('usernameOrEmail', 'Username or email is invalid', (value) => {
                return (
                    /^\S+@\S+\.\S+$/.test(value) || value.length > 0
                )
            }),
            password: Yup.string().min(6, "Password must be at least 6 characters").required("Password is required"),
        }),
        onSubmit: (values) => {
            console.log("Form Data: ", values)
        }
    })

    const handleGoogleLoginFailure = () => { 
        console.error("Google login failed")
    }
    const handleGoogleLoginSuccess = (response: any) => { 
        console.log("Google login success: ", response)
    }

  return (
    <div className="flex justify-center content-center items-center">

    <div className="container bg-white w-2/3 flex-col mx-4  border-2 rounded-2xl pt-12 my-5 justify-centera items-center content-center">
      <div className="flex justify-center h-full my-auto xl:gap-14 lg:justify-normal md:gap-5 draggable">
        <div className="flex items-center justify-center w-full lg:p-12">
          <div className="flex items-center xl:p-10">
            <form
              className="flex flex-col w-full h-full pb-6 text-center rounded-3xl"
              onSubmit={formik.handleSubmit}
            >
              <h3 className="mb-3 text-4xl font-extrabold text-dark-grey-900">Sign In</h3>
                <div className="flex my-4 justify-center content-center items-center">

              <GoogleLogin
                onSuccess={handleGoogleLoginSuccess}
                onError={handleGoogleLoginFailure}
                useOneTap
                theme="outline"
                text="signin_with"
                shape="circle"
                
              />
                </div>

              <div className="flex items-center mb-3">
                <hr className="h-0 border-b border-solid border-grey-500 grow" />
                <p className="mx-4 text-grey-600">or</p>
                <hr className="h-0 border-b border-solid border-grey-500 grow" />
              </div>
                <div>

              <label htmlFor="usernameOrEmail" className="mb-2 text-sm text-start text-grey-900">
                Username or Email*
              </label>
              <input
                id="usernameOrEmail"
                name="usernameOrEmail"
                type="text"
                placeholder="Enter username or email"
                className="flex items-center bg-gray-100 w-full px-5 py-4 mb-7 text-sm font-medium outline-none focus:bg-grey-400 placeholder:text-grey-700 bg-grey-200 text-dark-grey-900 rounded-2xl"
                onChange={formik.handleChange}
                onBlur={formik.handleBlur}
                value={formik.values.usernameOrEmail}
              />
              {formik.touched.usernameOrEmail && formik.errors.usernameOrEmail && (
                <div className="text-red-500 text-sm">{formik.errors.usernameOrEmail}</div>
              )}

              <label htmlFor="password" className="mb-2 text-sm text-start text-grey-900">
                Password *
              </label>
              <input
                id="password"
                name="password"
                type="password"
                placeholder="Enter a password"
                className="flex items-center bg-gray-100 w-full px-5 py-4 mb-5 text-sm font-medium outline-none focus:bg-grey-400 placeholder:text-grey-700 bg-grey-200 text-dark-grey-900 rounded-2xl"
                onChange={formik.handleChange}
                onBlur={formik.handleBlur}
                value={formik.values.password}
              />
              {formik.touched.password && formik.errors.password && (
                <div className="text-red-500 text-sm">{formik.errors.password}</div>
              )}

              <button
                type="submit"
                className="w-full px-6 py-5 mb-5 text-sm font-bold leading-none text-white transition duration-300 md:w-96 rounded-2xl hover:bg-purple-blue-600 focus:ring-4 bg-gradient-to-r from-indigo-400 to-indigo-600"
              >
                Sign In
              </button>
                </div>
              <p className="text-sm leading-relaxed text-grey-900">
                Not registered yet?{' '}
                <a href="javascript:void(0)" className="font-bold text-grey-700">
                  Create an Account
                </a>
              </p>
            </form>
          </div>
        </div>
      </div>
    </div>
    </div>
  )
}

export default LoginPage