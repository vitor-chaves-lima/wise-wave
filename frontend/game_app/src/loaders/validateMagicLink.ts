const validateMagicLink = async ({ request }: { request: Request }) => {
	const [,searchParams] = request.url.split("?");
	const challenge = new URLSearchParams(searchParams).get("challenge");

	const response = await fetch("https://tj6xftv7df.execute-api.sa-east-1.amazonaws.com/v1/auth/magic-link/validate", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({ challenge: challenge }),
	})

	if (response.status !== 200) {
		throw new Error("Não foi possível se comunicar com o servidor, gere um link de acesso novamente");
	}

	localStorage.setItem("tokenData", await response.text());

	return null;
}

export default validateMagicLink
