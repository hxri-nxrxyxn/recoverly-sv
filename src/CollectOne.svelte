<script>
  import { Link } from "svelte-routing";
  import NavClassic from "./NavClassic.svelte";
  import { Storage } from "@capacitor/storage";
  let firstname = $state("");
  let lastname = $state("");
  let date = $state("");
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
    const response = await fetch(
      `http://192.168.183.224:8080/api/v1/users/${id}`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      },
    );

    if (!response.ok) {
      const { error } = await response.json();
      alert(error);
      return;
    }

    const dataaparse = await response.json();

    const name = firstname + " " + lastname;
    const [day, month, year] = date.split("/");
    const birthdate = new Date(`${year}-${month}-${day}`);

    dataaparse.data.name = name;
    dataaparse.data.birthdate = birthdate;
    dataaparse.data.gender = gender;

    console.log(dataaparse.data);

    const res = await fetch(`http://192.168.183.224:8080/api/v1/users/${id}`, {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(dataaparse.data),
    });

    if (!res.ok) {
      const { error } = await res.json();
      alert(error);
      return;
    }

    const newdata = await res.json();
    console.log(newdata);

    location.href = "/collect-two";
  };
</script>

<main>
  <NavClassic />
  <h1>Let it be a new <span>Beginning</span></h1>
  <div class="form__box">
    <div class="form__link">
      <label>First Name</label>
      <input type="text" placeholder="Hari" required bind:value={firstname} />
    </div>
    <div class="form__link">
      <label>Last Name</label>
      <input type="text" placeholder="Narayan" required bind:value={lastname} />
    </div>
    <div class="form__link">
      <label>Date of Birth</label>
      <input type="date" placeholder="DD/MM/YY" required bind:value={date} />
    </div>
    <div class="form__link">
      <label>Gender</label>
      <input type="text" placeholder="Male" required bind:value={gender} />
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
