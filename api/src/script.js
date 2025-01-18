const checkName = async () => {
    try {
      const value = await Storage.get({ key: "id" });
      const id = value.value;
      if (!id) {
        hasToken = false;
        location.href = "/login";
      }
      const response = await fetch(
        `https://hardly-genuine-yeti.ngrok-free.app/api/v1/users/6`,
      );

      if (response.ok) {
        console.log("response ok");
        const data = await response.json();
      }
    } catch (error) {
      console.error("Error:", error);
    }
  };