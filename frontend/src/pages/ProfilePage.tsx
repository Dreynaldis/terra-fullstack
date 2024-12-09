import { useEffect, useState } from "react"
import axios from "axios"
import Cookies from "js-cookie"

const ProfilePage = () => {
    const [user, setUser] = useState<string>("")
    const [profile, setProfile] = useState<any>(null)

    const fetchUserData = async () => {
        try {
            const UsernameCookie = Cookies.get("username")
        setUser(UsernameCookie || "")
            const response = await axios.get(`${import.meta.env.VITE_API_BASE_URL}/users/${UsernameCookie}`)

            setProfile(response.data)
            // console.log(profile)

        } catch (error) {
            console.error("Failed to fetch user data: ", error)
        }
    }
    const formatDate = (dateString: string) => {
        const date = new Date(dateString)
        return date.toLocaleDateString() + " " + date.toLocaleTimeString()
    }
    useEffect(() => {
        fetchUserData()
    }, [])

    if (!profile) {
        return <div>Loading...</div>
    }
  return (
<div className="flex pt-10 container mx-auto content-center justify-center items-center">
      <div className="bg-p2 text-white p-6 justify-center items-center rounded shadow-md w-full max-w-lg">
        <h2 className="text-3xl font-semibold text-center mb-4">Welcome, {user}</h2>
        <div className="mb-4">
          <p><strong>Email:</strong> {profile.email}</p>
          <p><strong>Full Name:</strong> {profile.username}</p>
          <p><strong>You are logged in using:</strong> {profile.provider}</p>
        </div>
      </div>
    </div>
  )
}

export default ProfilePage