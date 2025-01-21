<script>
  import { Link } from "svelte-routing";
  import NavClassic from "./NavClassic.svelte";
  import { Storage } from "@capacitor/storage";
  let height = $state("");
  let weight = $state("");
  let severity = $state("");
  let conditions = $state("");

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

    dataaparse.data.height = Number(height);
    dataaparse.data.weight = Number(weight);
    dataaparse.data.severity = Number(severity);
    dataaparse.data.conditions = conditions;

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

    location.href = "/collect-three";
  };
</script>

<main>
  <NavClassic />
  <h1>Let it be a new <span>Beginning</span></h1>
  <div class="form__box">
    <div class="form__link">
      <label>Height</label>
      <input type="text" placeholder="CM" required bind:value={height} />
    </div>
    <div class="form__link">
      <label>Weight</label>
      <input type="password" placeholder="KG" required bind:value={weight} />
    </div>
    <div class="form__link">
      <label>Addiction Severity</label>
      <input type="range" min="1" max="100" bind:value={severity} />
    </div>
    <div class="form__link">
      <label>Existing Medical Conditions</label>
      <input
        type="text"
        placeholder="If any"
        required
        bind:value={conditions}
      />
    </div>
  </div>
  <div class="btn__box">
    <div class="progress">
      <div class="progressi">whoa lion!</div>
    </div>
    <button class="solid" onclick={collecttwo}>NEXT</button>
  </div>
</main>

<style>
  .progressi {
    width: 60%;
  }
</style>
