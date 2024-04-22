from django.db import models
from django.contrib.auth.models import User

class Product(models.Model): #creating models for products i.e for medicine
    name = models.CharField(max_length=100)
    price = models.DecimalField(max_digits=10, decimal_places=2)
    stock = models.PositiveIntegerField()
    status = models.BooleanField(default=True)
    image = models.ImageField(upload_to='product_images/', blank=True, null=True)

    def is_available(self):
        return "Available" if self.status else "Out of stock" #alias for status
    
    def __str__(self):
        return self.name #setting up the object name

class Message(models.Model):
    sender=models.ForeignKey(User, related_name='sent_messages', on_delete=models.CASCADE)
    recipient=models.ForeignKey(User, related_name='recieve_messages', on_delete=models.CASCADE)
    content=models.TextField()
    timestamp=models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return f"From {self.sender.username} to {self.recipient.username}"

class RequestBlood(models.Model):
    name = models.CharField(max_length=100)
    address = models.CharField(max_length=255)
    gender = models.CharField(max_length=10)
    dob = models.DateField()
    age = models.IntegerField()
    blood_group = models.CharField(max_length=5)
    phone = models.CharField(max_length=15)
    email = models.EmailField()

    def __str__(self):
        return f"bloodRequest by {self.name}"
    
class DonateBlood(models.Model):
    name = models.CharField(max_length=100)
    address = models.CharField(max_length=255)
    gender = models.CharField(max_length=10)
    dob = models.DateField()
    age = models.IntegerField()
    blood_group = models.CharField(max_length=5)
    phone = models.CharField(max_length=15)
    email = models.EmailField()

    def __str__(self):
        return f"blood donation by {self.name}"
    
class CartItem(models.Model):
    user = models.ForeignKey(User, on_delete=models.CASCADE)
    product = models.ForeignKey(Product, on_delete=models.CASCADE)
    quantity = models.PositiveIntegerField(default=1)

    def subtotal(self):
        return self.product.price * self.quantity
    