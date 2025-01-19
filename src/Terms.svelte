<script>
  import { Link } from "svelte-routing";
</script>

<main>
  <h1>Let it be a new <span>Beginning</span></h1>

  <div class="text__box">
    <h2>Terms & Conditions</h2>
    <p>
      By using Recoverly, you agree to these Terms and Conditions. The app
      provides resources and support for those overcoming addiction but is not a
      substitute for professional medical advice. Always consult healthcare
      professionals for personalized guidance. By accessing the app, you confirm
      you are at least 15 years old and will use the app legally and
      appropriately. You agree not to engage in harmful, illegal activities or
      disrupt the app’s functionality. We respect your privacy and handle your
      data according to our <b>Privacy Policy</b>. You are responsible for
      securing your account. Recoverly is not liable for any consequences from
      using the app. We reserve the right to modify these terms or suspend
      access for violations.
      <br />
      <br />
      <b> Privacy Policy </b>
      <br />
      Your privacy is important to us. Recoverly collects personal information only
      to enhance your user experience. We do not share your data with third parties
      without your consent, except as required by law. For more details, please refer
      to the full Privacy Policy within the app.
    </p>
  </div>

  <div class="btn__box">
    <Link to="/collect-one">
      <button class="solid">CONTINUE</button>
    </Link>
  </div>
</main>

<style>
  h2 {
    margin-bottom: 1rem;
  }
</style>
