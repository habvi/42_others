import smtplib
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart
from email.mime.application import MIMEApplication
from dotenv import load_dotenv
import os

# from: SMTP
host = "smtp.gmail.com"
port = 587

# email content
subject = "event april 2023 / ex12 test!"
message = "hello- hello- from python"

# file attachment
attachment_path = "42.png"

def create_MIME():
    from_email = os.environ["FROM_EMAIL"]
    to_email = os.environ["TO_EMAIL"]

    mime = MIMEMultipart()
    mime["Subject"] = subject
    mime["From"] = from_email
    mime["To"] = to_email
    mime["Bcc"] = from_email
    mime.attach(MIMEText(message))

    with open(attachment_path, "rb") as f:
        attachment_file = MIMEApplication(f.read())
        attachment_file.add_header(
            "Content-Disposition",
            "attachment",
            filename=attachment_path
        )
        mime.attach(attachment_file)
    return mime


def send_mail(mime):
    account = os.environ["FROM_EMAIL"]
    password = os.environ["FROM_PASSWORD"]

    try:
        server = smtplib.SMTP(host, port)
        server.starttls()
        server.login(account, password)
        server.send_message(mime)
        server.quit()
    except Exception as E:
        print("=== Error ===")
        print("Error:", str(E))
        print("type: ", str(type(E)))
        print("args:", str(E.args))
        print("message:", E.message)


def main():
    load_dotenv()
    mime = create_MIME()
    send_mail(mime)


if __name__ == '__main__':
    main()