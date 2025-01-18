<script>
  import { Link } from "svelte-routing";
  import NavClassic from "./NavClassic.svelte";
  import { Storage } from "@capacitor/storage";
  let firstname = $state("");
  let lastname = $state("");
  let birthdate = $state("");
  let gender = $state("");

  const checkUser = async () => {
    const value = await Storage.get({ key: "id" });
    const id = value.value;
    console.log(id);
    if (!id) {
      location.href = "/login";
    }
  };
  checkUser();

  const collecttwo = async () => {
    const value = await Storage.get({ key: "id" });
    const id = value.value;
    console.log(value);
    const response = await fetch(`http://localhost:8080/api/v1/users/${id}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      const { error } = await response.json();
      alert(error);
      return;
    }

    const name = firstname + " " + lastname;

    const dataaparse = {
      name: name,
      birthdate: birthdate,
      gender: gender,
    };

    const res = await fetch(`http://localhost:8080/api/v1/users/${id}`, {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(dataaparse),
    });

    if (!res.ok) {
      const { message } = await res.json();
      alert(message);
      return;
    }

    location.href = "/collect-two";
  };
</script>

<main>
  <NavClassic />
  <h1>Let it be a new <span>Beginning</span></h1>
  <div class="form__box">
    <div class="form__link">
      <label>First Name</label>
      <input type="text" placeholder="Hari" required />
    </div>
    <div class="form__link">
      <label>Last Name</label>
      <input type="password" placeholder="Narayan" required />
    </div>
    <div class="form__link">
      <label>Date of Birth</label>
      <input type="date" placeholder="DD/MM/YY" required />
    </div>
    <div class="form__link">
      <label>Gender</label>
      <input type="password" placeholder="Male" required />
    </div>
  </div>
  <div class="btn__box">
    <div class="progress">
      <div class="progressi">whoa lion!</div>
    </div>
    <button class="solid" onclick={collecttwo}>CONTINUE</button>
  </div>
</main>

<style>
  .progressi {
    width: 30%;
  }
</style>
