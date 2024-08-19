const sendMagicLinkAction = async ({ request }: { request: Request }) => {
    const data = await request.formData()
    const email = data.get("email")

    await fetch("https://tj6xftv7df.execute-api.sa-east-1.amazonaws.com/v1/auth/magic-link", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ email: email }),
    })

    return null;
}

export default sendMagicLinkAction